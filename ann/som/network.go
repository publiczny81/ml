package som

import (
	"github.com/publiczny81/ml/ann/neuron"
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/metrics"
	"github.com/publiczny81/ml/utils"
	"github.com/publiczny81/ml/utils/slices"
	"math"
)

type Metrics struct {
	Name string
	metrics.Metrics
}

var (
	defaultNetworkConfig = config{
		Metrics:  metrics.Euclidean,
		Topology: TopologyLinear,
		Rand:     utils.Rand,
	}
)

// Rand is a contract for random number generator used for randomizing weights
type Rand interface {
	Float64() float64
}

// Point represents a point in n-dimensional space
type Point []float64

// Neuron represents a neuron in Self-Organizing Map network
type Neuron struct {
	Point
	*neuron.Neuron[float64]
}

func NewNeuron(point Point, n *neuron.Neuron[float64]) *Neuron {
	return &Neuron{
		Point:  point,
		Neuron: n,
	}
}

// config contains configuration of the network
type config struct {
	// Features is input vector size
	Features int
	// Metrics which are used for calculation of distance between input vector and weights of the neuron
	Metrics  string
	// Rand is a number generator used for initialization of neuron's w
	Rand     Rand
	// Shape represents a shape of the network
	Shape    []int
	// Topology represents topology of the network
	Topology int
	// Weights contains the weights of the neurons
	Weights  []float64
}

type Option func(options *config) error

func WithWeights(weights []float64) Option {
	return func(options *config) error {
		if len(weights) == 0 {
			return errors.WithMessage(errors.InvalidParameterValueError, "len(weights)=0")
		}
		if len(weights) != slices.Aggregate(options.Shape, options.Features, func(acc int, factor int) int {
			acc *= factor
			return acc
		}) {
			return errors.WithMessagef(errors.InvalidParameterValueError, "incompatible shape=%v and len(weights)=%d", options.Shape, len(weights))
		}
		options.Weights = weights
		return nil
	}
}

// WithMetrics sets the metrics function used to calculate the distance between features vector and neuron weights
// The default is Euclidean distance
func WithMetrics(metricName string) Option {
	return func(options *config) error {
		if _, found := metrics.Get(metricName); !found {
			return errors.WithMessagef(errors.InvalidParameterValueError, "metrics=%s", metricName)
		}
		options.Metrics = metricName
		return nil
	}
}

func WithTopology(topology int) Option {
	return func(options *config) error {
		if topology < TopologyLinear || topology > TopologyHexagonal {
			return errors.WithMessagef(errors.InvalidParameterValueError, "topology=%d", topology)
		}
		options.Topology = topology
		return nil
	}
}

func WithRand(rand Rand) Option {
	return func(options *config) error {
		if rand == nil {
			return errors.WithMessage(errors.InvalidParameterValueError, "rand=nil")
		}
		options.Rand = rand
		return nil
	}
}

// Network represents a Self-Organizing Map network
type Network struct {
	config
	Neurons []*Neuron
}

func New(features int, shape []int, opts ...Option) (n *Network, err error) {
	var (
		cfg = defaultNetworkConfig
	)

	cfg.Features = features
	cfg.Shape = shape

	for _, o := range opts {
		if err = o(&cfg); err != nil {
			return
		}
	}

	n = &Network{
		config: cfg,
	}
	return
}

func validateConfig(config *config) (err error) {
	if config.Features <= 0 {
		err = errors.WithMessagef(errors.InvalidParameterValueError, "features=%d", config.Features)
		return
	}
	if len(config.Shape) == 0 {
		err = errors.WithMessagef(errors.InvalidParameterValueError, "shape=%v", config.Shape)
		return
	}
	// adjust shape to topology
	switch config.Topology {
	case TopologyLinear:
		config.Shape = config.Shape[:1]
	case TopologyRectangular, TopologyHexagonal:
		if len(config.Shape) < 2 {
			return errors.WithMessagef(errors.InvalidParameterValueError, "shape=%v", config.Shape)
		}
		config.Shape = config.Shape[:2]
	default:
	}

	// validate shape
	for i := range config.Shape {
		if config.Shape[i] <= 0 {
			err = errors.WithMessagef(errors.InvalidParameterValueError, "shape=%v", config.Shape)
			return
		}
	}

	// validate weights
	if len(config.Weights) > 0 && len(config.Weights) < slices.Aggregate(config.Shape, config.Features, func(acc int, factor int) int {
		acc *= factor
		return acc
	}) {
		err = errors.WithMessagef(errors.InvalidParameterValueError, "incompatible shape=%v and len(weights)=%d", config.Shape, len(config.Weights))
		return
	}

	if config.Rand == nil {
		config.Rand = defaultNetworkConfig.Rand
	}

	return nil
}

func (net *Network) Init(opts ...Option) (err error) {
	for _, o := range opts {
		if err = o(&net.config); err != nil {
			return err
		}
	}

	if err = validateConfig(&net.config); err != nil {
		return
	}

	net.initWeights()

	metric, _ := metrics.Get(net.config.Metrics)

	start, end := 0, net.Features

	for point := range NewGenerator(net.config.Topology, net.config.Shape...) {
		net.Neurons = append(net.Neurons, &Neuron{
			Point:  point,
			Neuron: neuron.New(metric.Function, net.config.Weights[start:end]),
		})
		start = end
		end += net.Features
	}

	return nil
}

func (net *Network) randomize() {
	for i := range net.Weights {
		net.Weights[i] = net.Rand.Float64()
	}
}

func (net *Network) initWeights() {
	if len(net.Weights) == 0 {
		count := slices.Aggregate(net.config.Shape, net.config.Features, func(acc int, factor int) int {
			acc *= factor
			return acc
		})
		net.Weights = make([]float64, count)
		net.randomize()
	}
}

func (net *Network) BestMatchingUnit(input []float64) (bmu Point) {
	var (
		minDistance = math.MaxFloat64
	)

	for i := range net.Neurons {
		var distance = net.Neurons[i].Activate(input)
		if distance < minDistance {
			minDistance = distance
			bmu = net.Neurons[i].Point
		}
	}

	return
}

package som

import (
	"github.com/pkg/errors"
	"github.com/publiczny81/ml/array"
	"github.com/publiczny81/ml/metrics"
	"github.com/publiczny81/ml/types"
	"github.com/publiczny81/ml/utils/slice"
	"math"
)

var (
	InvalidParameterError = errors.New("invalid parameter")
)

var (
	defaultNetworkOptions = networkOptions{
		Metrics: metrics.EuclideanDistance,
	}
)

type Rand interface {
	Float64() float64
}

type neuron []float64

func (n neuron) Process(metrics types.MetricsFunc, input []float64) float64 {
	return metrics(input, n)
}

type networkOptions struct {
	Metrics  types.MetricsFunc
	Weights  []float64
	Features int
	Dim      []int
	Rand     Rand
}

type NetworkOption func(options *networkOptions)

func WithWeights(weights []float64) NetworkOption {
	return func(options *networkOptions) {
		options.Weights = weights
	}
}

func WithMetrics(metrics types.MetricsFunc) NetworkOption {
	return func(options *networkOptions) {
		options.Metrics = metrics
	}
}

func WithRand(rand Rand) NetworkOption {
	return func(options *networkOptions) {
		options.Rand = rand
	}
}

type Network struct {
	Metrics types.MetricsFunc
	Weights []float64
	Neurons *array.Array[neuron]
}

func New(features int, dim []int, opts ...NetworkOption) (n *Network, err error) {
	var (
		options = defaultNetworkOptions
	)

	options.Features = features
	options.Dim = dim

	for _, o := range opts {
		o(&options)
	}

	if err = validateAndInitNetworkOptions(&options); err != nil {
		return
	}

	var (
		start, end = 0, options.Features
	)

	n = &Network{
		Metrics: options.Metrics,
		Weights: options.Weights,
		Neurons: array.NewBuilder[neuron](dim...).
			WithInitFunc(func(idx int) (n neuron) {
				n = options.Weights[start:end]
				start = end
				end += features
				return
			}).Build(),
	}

	return
}

func (net *Network) Randomize(rand Rand) {
	for i := range net.Weights {
		net.Weights[i] = rand.Float64()
	}
}

func (net *Network) Dim() []int {
	return net.Neurons.Dim()
}

func (net *Network) Features() int {
	return len(net.Weights) / net.Neurons.Size()
}

func (net *Network) BestMatchingUnit(input []float64) []int {
	var (
		bmu         int
		minDistance = math.MaxFloat64
	)

	net.Neurons.IterateWithIndex(func(idx int, neuron neuron) bool {
		var distance = neuron.Process(net.Metrics, input)
		if distance < minDistance {
			minDistance = distance
			bmu = idx
		}
		return true
	})

	return net.Neurons.Position(bmu)
}

func validateAndInitNetworkOptions(options *networkOptions) (err error) {
	if options.Features == 0 {
		err = errors.WithMessage(InvalidParameterError, "missing features number")
		return
	}
	if len(options.Dim) == 0 {
		err = errors.WithMessage(InvalidParameterError, "missing dimensions")
		return
	}
	var length = slice.Aggregate(options.Dim, options.Features, func(i int, i2 int) int {
		return i * i2
	})
	if len(options.Weights) > 0 && len(options.Weights) < length {
		err = errors.WithMessage(InvalidParameterError, "incompatible dim and weights")
		return
	}
	if len(options.Weights) == 0 && (options.Rand == nil) {
		err = errors.WithMessage(InvalidParameterError, "missing rand function")
		return
	}
	if len(options.Weights) == 0 {
		options.Weights = make([]float64, slice.Aggregate(options.Dim, options.Features, func(i int, i2 int) int {
			return i * i2
		}))
		slice.Apply(options.Weights, func(f float64) float64 {
			return options.Rand.Float64()
		})
	}
	if options.Metrics == nil {
		err = errors.WithMessage(InvalidParameterError, "metrics")
		return
	}
	return
}

package codecs

import (
	"bytes"
	"encoding/json"
	"github.com/publiczny81/ml/ann/som"
	"github.com/publiczny81/ml/ann/som/model"
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/metrics"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestEncoder(t *testing.T) {
	//var (
	//	input         = []float64{1, 2, 3, 4}
	//	buf           = bytes.NewBuffer(nil)
	//	enc           = NewEncoder(buf)
	//	expected, err = som.New(4, []int{10, 10}, som.WithTopology(som.TopologyHexagonal))
	//	actual        = new(som.Network)
	//)
	//assert.NoError(t, err)
	//assert.NoError(t, expected.Init())
	//assert.NoError(t, enc.Encode(expected))
	//assert.NoError(t, Decode(buf.Bytes(), actual))
	//assert.NoError(t, actual.Init())
	//assert.Equal(t, expected.BestMatchingUnit(input), actual.BestMatchingUnit(input))

	var tests = []struct {
		Name     string
		Input    any
		Expected []byte
		Error    error
	}{
		{
			Name:  "When passed value is nil then return the error",
			Input: (*som.Network)(nil),
			Error: errors.InvalidParameterValueError,
		},
		{
			Name:  "When passed value is neither *som.Network nor som.Network then return the error",
			Input: struct{}{},
			Error: errors.InvalidParameterValueError,
		},
		{
			Name: "When passed value is som.Network then encode the network",
			Input: func() (n som.Network) {
				network, _ := som.New(1, []int{2}, som.WithWeights([]float64{1, 2}))
				n = *network
				return
			}(),
			Expected: func() []byte {
				var buffer = new(bytes.Buffer)
				_ = json.NewEncoder(buffer).Encode(&model.Network{
					Features: 1,
					Metrics:  metrics.Euclidean,
					Shape:    []int{2},
					Topology: som.TopologyLinear,
					Weights:  []float64{1, 2},
				})
				return buffer.Bytes()
			}(),
		},
		{
			Name: "When passed value is *som.Network then encode the network",
			Input: func() (n *som.Network) {
				n, _ = som.New(1, []int{2}, som.WithWeights([]float64{1, 2}))
				return
			}(),
			Expected: func() []byte {
				var buffer = new(bytes.Buffer)
				_ = json.NewEncoder(buffer).Encode(&model.Network{
					Features: 1,
					Metrics:  metrics.Euclidean,
					Shape:    []int{2},
					Topology: som.TopologyLinear,
					Weights:  []float64{1, 2},
				})
				return buffer.Bytes()
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual = new(bytes.Buffer)
				err    = NewEncoder(actual).Encode(test.Input)
			)
			if test.Error != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, test.Error.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, actual.Bytes())
		})
	}
}

func TestDecoder(t *testing.T) {
	var (
		tests = []struct {
			Name     string
			Buffer   []byte
			Input    any
			Expected *som.Network
			Error    error
		}{
			{
				Name:  "When invalid data given then return the error",
				Error: io.EOF,
			},
			{
				Name: "When valid data given then decode the network",
				Buffer: func() []byte {
					buffer := bytes.NewBuffer(nil)
					net, _ := som.New(1, []int{2}, som.WithWeights([]float64{1, 2}))
					_ = NewEncoder(buffer).Encode(net)
					return buffer.Bytes()
				}(),
				Expected: func() *som.Network {
					net, _ := som.New(1, []int{2}, som.WithWeights([]float64{1, 2}))
					return net
				}(),
			},
		}
	)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var (
				actual = new(som.Network)
				err    = NewDecoder(bytes.NewReader(test.Buffer)).Decode(actual)
			)
			if test.Error != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, test.Error.Error())
				return
			}
			assert.NoError(t, err)
			assert.Condition(t, func() (success bool) {
				if !assert.Equal(t, test.Expected.Features, actual.Features) {
					return
				}
				if !assert.Equal(t, test.Expected.Shape, actual.Shape) {
					return
				}
				if !assert.Equal(t, test.Expected.Metrics, actual.Metrics) {
					return
				}
				if !assert.Equal(t, test.Expected.Topology, actual.Topology) {
					return
				}
				if !assert.Equal(t, test.Expected.Weights, actual.Weights) {
					return
				}
				success = true
				return
			})
		})
	}
}

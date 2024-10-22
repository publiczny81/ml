package som

import (
	"encoding/json"
	"github.com/publiczny81/ml/ann/som"
	"github.com/publiczny81/ml/errors"
	"io"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer: writer,
	}
}

func (enc *Encoder) Encode(v any) (err error) {
	switch value := v.(type) {
	case *som.Network:
		return enc.encode(value)
	case som.Network:
		return enc.encode(&value)
	default:
		err = errors.WithMessagef(errors.InvalidParameterValueError, "v is neither *som.Network nor som.Network")
		return
	}
}

func (enc *Encoder) encode(network *som.Network) error {
	if network == nil {
		return errors.WithMessage(errors.InvalidParameterValueError, "network is nil")
	}
	var net = &Network{
		Features: network.Features,
		Metrics:  network.Metrics,
		Shape:    network.Shape,
		Topology: network.Topology,
		Weights:  network.Weights,
	}

	return json.NewEncoder(enc.writer).Encode(net)
}

type Decoder struct {
	reader io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		reader: reader,
	}
}

func (dec *Decoder) Decode(v any) (err error) {
	switch value := v.(type) {
	case *som.Network:
		return dec.decode(value)
	default:
		err = errors.WithMessage(errors.InvalidParameterError, "v must be *som.Network")
		return
	}
}

func (dec *Decoder) decode(network *som.Network) (err error) {
	if network == nil {
		err = errors.WithMessage(errors.InvalidParameterValueError, "network is nil")
		return
	}
	var net = new(Network)
	if err = json.NewDecoder(dec.reader).Decode(net); err != nil {
		return
	}
	network.Features = net.Features
	network.Shape = net.Shape
	network.Metrics = net.Metrics
	network.Topology = net.Topology
	network.Weights = net.Weights

	return
}

func Decode(buffer []byte, network *som.Network) (err error) {
	if network == nil {
		err = errors.WithMessage(errors.InvalidParameterValueError, "network is nil")
		return
	}
	var net = new(Network)
	if err = json.Unmarshal(buffer, net); err != nil {
		return
	}
	network.Features = net.Features
	network.Shape = net.Shape
	network.Metrics = net.Metrics
	network.Topology = net.Topology
	network.Weights = net.Weights

	return
}

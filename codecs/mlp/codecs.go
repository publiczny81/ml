package mlp

import (
	"encoding/json"
	"github.com/publiczny81/ml/ann/mlp"
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
	case *mlp.Network:
		return enc.encode(value)
	case mlp.Network:
		return enc.encode(&value)
	default:
		err = errors.WithMessagef(errors.InvalidParameterValueError, "v is neither *mlp.Network nor mlp.Network")
		return
	}
}

func (enc *Encoder) encode(network *mlp.Network) error {
	if network == nil {
		return errors.WithMessage(errors.InvalidParameterValueError, "network is nil")
	}
	var net = new(Network)
	net.Input = network.Input
	net.Weights = network.Options.Weights
	for _, l := range network.Options.Layers {
		var layer = new(LayerSpec)
		layer.Neurons = l.Neurons

		net.Layers = append(net.Layers, LayerSpec{
			Neurons:    l.Neurons,
			Activation: l.Activation,
		})
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

func (d *Decoder) Decode(v any) (err error) {
	switch value := v.(type) {
	case *mlp.Network:
		return d.decode(value)
	default:
		err = errors.WithMessagef(errors.InvalidParameterValueError, "v must be *mlp.Network")
		return
	}
}

func (d *Decoder) decode(network *mlp.Network) error {
	if network == nil {
		return errors.WithMessage(errors.InvalidParameterValueError, "network is nil")
	}
	var net = new(Network)
	if err := json.NewDecoder(d.reader).Decode(net); err != nil {
		return err
	}
	network.Options.Input = net.Input
	for _, l := range net.Layers {
		network.Options.Layers = append(network.Options.Layers, mlp.LayerSpec{
			Neurons:    l.Neurons,
			Activation: l.Activation,
		})
	}
	network.Options.Weights = net.Weights
	return nil
}

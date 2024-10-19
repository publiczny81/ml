package mlp

//
//import (
//	"github.com/pkg/errors"
//	"github.com/publiczny81/ml/ann/neuron"
//	"github.com/publiczny81/ml/calculus/types"
//	errors2 "github.com/publiczny81/ml/errors"
//	"github.com/publiczny81/ml/utils/slices"
//)
//
//type sampler[T types.Float] interface {
//	Next() [][]T
//	HasNext() bool
//	Reset()
//}
//
//type learningRate[T types.Float] interface {
//	LearningRate(epoch int) T
//}
//
//type lossFunction[T types.Float] func([]T, []T) ([]T, T)
//
//type BackPropagationTrainer[T types.Float] struct {
//	sampler[T]
//	learningRate[T]
//	lossFunction[T]
//	dActivateFunc func(T) T
//	epochs        int
//	stopCondition func(T) bool
//	errors        [][]T
//}
//
//type NeuronWithError[T types.Float] struct {
//	*neuron.Neuron[T]
//	Error T
//}
//
//func (t *BackPropagationTrainer[T]) Train(net *Network[T]) (err T, epoch int, trained bool) {
//	if net == nil {
//		panic(errors.WithMessage(errors2.InvalidParameterError, "BackPropagationTrainer.Train: net is nil"))
//	}
//	t.init(net)
//	for epoch = range t.epochs {
//		err = t.train(net, epoch)
//		if trained = t.stopCondition(err); trained {
//			return
//		}
//	}
//	return
//}
//
//func (t *BackPropagationTrainer[T]) init(net *Network[T]) {
//	t.errors = make([][]T, 0, len(net.Neurons))
//	slices.Reverse(net.Neurons, func(l Layer[T]) bool {
//		t.errors = append(t.errors, make([]T, len(l)))
//		return true
//	})
//}
//
//func (t *BackPropagationTrainer[T]) train(net *Network[T], epoch int) (err T) {
//	var lr = t.learningRate.LearningRate(epoch)
//
//	t.sampler.Reset()
//
//	for t.sampler.HasNext() {
//		var (
//			sample         = t.sampler.Next()
//			partials, loss = t.forwardStep(net, sample)
//		)
//		t.backwardStep(net, partials)
//		t.nudgeNeurons(net, lr)
//		err += loss
//	}
//	return
//}
//
//func (t *BackPropagationTrainer[T]) forwardStep(net *Network[T], sample [][]T) (partials []T, loss T) {
//	partials, loss = t.lossFunction(sample[1], net.Process(sample[0]))
//	return
//}
//
//func (t *BackPropagationTrainer[T]) backwardStep(net *Network[T], partials []T) {
//	t.errors[0] = partials
//	//slice.Reverse(net.Neurons, func(l Layer[T]) bool {
//	//	slice.Reverse(l, func(n *NeuronWithError[T]) bool {
//	//		t.errors[0] = slice.Apply(t.errors[0], func(i int, e T) T {
//	//			return n.Error * n.ActivationFunc.Derivative(n.Inputs[i])
//	//		})
//	//		return true
//	//	})
//	//	return true
//	//})
//	panic("implement me")
//}
//
//func (t *BackPropagationTrainer[T]) nudgeNeurons(net *Network[T], lr T) {
//	panic("implement me")
//}

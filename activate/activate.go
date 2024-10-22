package activate

import (
	"github.com/publiczny81/ml/errors"
	"github.com/publiczny81/ml/functions"
	"strconv"
	"strings"
	"sync"
)

const (
	separator = "@"
)
const (
	Linear    = "linear"
	Sigmoid   = "sigmoid"
	Rectifier = "rectifier"
)

type Activate struct {
	Name       string
	Function   func(float64) float64
	Derivative func(float64) float64
}

type factory func(...any) Activate

func Get(name string, params ...any) (a Activate, found bool) {
	registerOnce.Do(func() {
		register = map[string]factory{
			Linear:    GetLinear,
			Sigmoid:   GetSigmoid,
			Rectifier: GetRectifier,
		}
	})

	var (
		names = strings.Split(name, separator)
		f     factory
	)
	if f, found = register[names[0]]; !found {
		return
	}

	for _, sub := range names[1:] {
		params = append(params, sub)
	}

	a = f(params...)
	return
}

var (
	registerOnce sync.Once
	register     = map[string]factory{}
)

func GetSigmoid(_ ...any) (a Activate) {
	a.Name = Sigmoid
	a.Function = functions.Sigmoid
	a.Derivative = func(value float64) float64 {
		var ret = functions.Sigmoid(value)
		return ret * (1.0 - ret)
	}
	return
}

func GetRectifier(params ...any) (a Activate) {
	if len(params) == 0 {
		panic(errors.WithStack(errors.New("activate.GetRectifier: missing parameter")))
		return
	}
	var value, err = float64FromAny(params[0])
	if err != nil {
		panic(errors.WithStack(err))
		return
	}
	a.Name = "rectifier@" + strconv.FormatFloat(value, 'f', -1, 64)
	a.Function = functions.ParametricRectifier(value)
	a.Derivative = functions.DerivativeParametricRectifier(value)
	return
}

func GetLinear(params ...any) (f Activate) {
	var (
		a, b float64
		err  error
	)
	switch len(params) {
	case 0:
		a, b = 1.0, 0.0
	case 1:
		if a, err = float64FromAny(params[0]); err != nil {
			panic(errors.WithStack(err))
			return
		}
		b = 0.0
	default:
		if a, err = float64FromAny(params[0]); err != nil {
			panic(errors.WithStack(err))
			return
		}
		if b, err = float64FromAny(params[1]); err != nil {
			panic(errors.WithStack(err))
			return
		}
	}

	f.Name = "linear@" + strconv.FormatFloat(a, 'f', -1, 64) + separator + strconv.FormatFloat(b, 'f', -1, 64)
	f.Function = functions.Linear(a, b)
	f.Derivative = functions.DerivativeLinear(a)
	return
}

func float64FromAny(value any) (float64, error) {
	switch val := value.(type) {
	case float64:
		return val, nil
	case int8, int16, int32, int64, int, float32:
		return val.(float64), nil
	case string:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, errors.New("activate.float64FromAny: invalid type")
	}
}

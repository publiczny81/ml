package pool

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/utils/slices"
	"sync"
	"sync/atomic"
)

var (
	latestVectorSizeInt8 atomic.Int32
	vectorPoolInt8       = sync.Pool{
		New: func() interface{} {
			return make([]int8, 0, int(latestVectorSizeInt8.Load()))
		},
	}
)

var (
	latestVectorSizeInt16 atomic.Int32
	vectorPoolInt16       = sync.Pool{
		New: func() interface{} {
			return make([]int16, 0, int(latestVectorSizeInt16.Load()))
		},
	}
)

var (
	latestVectorSizeInt32 atomic.Int32
	vectorPoolInt32       = sync.Pool{
		New: func() interface{} {
			return make([]int32, 0, int(latestVectorSizeInt32.Load()))
		},
	}
)

var (
	latestVectorSizeInt64 atomic.Int32
	vectorPoolInt64       = sync.Pool{
		New: func() interface{} {
			return make([]int, 0, int(latestVectorSizeInt64.Load()))
		},
	}
)

var (
	latestVectorSizeInt atomic.Int32
	vectorPoolInt       = sync.Pool{
		New: func() interface{} {
			return make([]int, 0, int(latestVectorSizeInt.Load()))
		},
	}
)

var (
	latestVectorSizeFloat32 atomic.Int32
	vectorPoolFloat32       = sync.Pool{
		New: func() interface{} {
			return make([]float32, 0, latestVectorSizeFloat32.Load())
		},
	}
)

var (
	latestVectorSizeFloat64 atomic.Int32

	vectorPoolFloat64 = sync.Pool{
		New: func() interface{} {
			return make([]float64, 0, latestVectorSizeFloat64.Load())
		},
	}
)

func Get[S ~[]T, T types.Real](size int) (v S) {
	if size < 0 {
		return
	}
	switch any(v).(type) {
	case []int8, types.V[int8]:
		latestVectorSizeInt8.Store(int32(size))
		v, _ = vectorPoolInt8.Get().(S)
	case []int16, types.V[int16]:
		latestVectorSizeInt16.Store(int32(size))
		v, _ = vectorPoolInt16.Get().(S)
	case []int32, types.V[int32]:
		latestVectorSizeInt32.Store(int32(size))
		v, _ = vectorPoolInt32.Get().(S)
	case []int64, types.V[int64]:
		latestVectorSizeInt64.Store(int32(size))
		v, _ = vectorPoolInt64.Get().(S)
	case []int, types.V[int]:
		latestVectorSizeInt.Store(int32(size))
		v, _ = vectorPoolInt.Get().(S)
	case []float32, types.V[float32]:
		latestVectorSizeFloat32.Store(int32(size))
		v, _ = vectorPoolFloat32.Get().(S)
	case []float64, types.V[float64]:
		latestVectorSizeFloat64.Store(int32(size))
		v, _ = vectorPoolFloat64.Get().(S)
	}
	return slices.Resize(v, size)
}

func Put[S ~[]T, T types.Real](v S) {
	if v == nil {
		return
	}
	clear(v[:cap(v)])

	switch any(v).(type) {
	case []int8:
		vectorPoolInt8.Put(v)
	case []int16:
		vectorPoolInt16.Put(v)
	case []int32:
		vectorPoolInt32.Put(v)
	case []int64:
		vectorPoolInt64.Put(v)
	case []int:
		vectorPoolInt.Put(v)
	case []float32:
		vectorPoolFloat32.Put(v)
	case []float64:
		vectorPoolFloat64.Put(v)
	}
}

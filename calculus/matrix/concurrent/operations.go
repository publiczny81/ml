package concurrent

import (
	"github.com/publiczny81/ml/calculus/types"
	"github.com/publiczny81/ml/calculus/vector"
	"github.com/publiczny81/ml/calculus/vector/pool"
	"runtime"
	"sync"
)

var (
	maxThreads = runtime.NumCPU()*2 - 1
)

type task func()

func worker(wg *sync.WaitGroup, tasks <-chan task) {
	for t := range tasks {
		t()
	}
	wg.Done()
}

func Column[M ~[][]T, T types.Real](m M, j int) (result []T) {
	result = pool.Get[[]T](len(m))
	for i := range result {
		result[i] = m[i][j]
	}
	return
}

func ForEachColumn[T types.Real](f func(j int, c []T) []T) types.MOperation[T] {
	return func(m types.M[T]) {
		var (
			wg    sync.WaitGroup
			size  = len(m) / maxThreads
			tasks = make(chan task, size)
		)

		for range maxThreads {
			wg.Add(1)
			go worker(&wg, tasks)
		}

		go func() {
			for i := range len(m[0]) {
				tasks <- func() {
					c := Column(m, i)
					result := f(i, c)
					for j, v := range result {
						m[j][i] = v
					}
					pool.Put(f(i, c))
				}
			}
			close(tasks)
		}()
		wg.Wait()
	}
}

func ForEachRow[T types.Real](f func(i int, r []T) []T) types.MOperation[T] {
	return func(m types.M[T]) {
		var (
			wg   sync.WaitGroup
			size = len(m) / maxThreads
		)

		var ch = make(chan task, size)

		for range maxThreads {
			wg.Add(1)
			go func() {
				for s := range ch {
					s()
				}
				wg.Done()
			}()
		}

		go func() {
			for i := 0; i < len(m); i++ {
				ch <- func() {
					m[i] = f(i, m[i])
				}
			}
			close(ch)
		}()

		wg.Wait()
	}
}

func ForEach[T types.Real](f func(i, j int, t T) T) types.MOperation[T] {
	return func(m types.M[T]) {
		var (
			wg   sync.WaitGroup
			size = len(m) / maxThreads
		)

		var tasks = make(chan task, size)

		for range maxThreads {
			wg.Add(1)
			go worker(&wg, tasks)
		}

		go func() {
			for i := 0; i < len(m); i++ {
				for j := 0; j < len(m[i]); j++ {
					tasks <- func() {
						m[i][j] = f(i, j, m[i][j])
					}
				}
			}
			close(tasks)
		}()

		wg.Wait()
	}
}

func Add[R ~[][]T, T types.Real](m R) types.MOperation[T] {
	return ForEachRow[T](func(i int, r []T) []T {
		return vector.Add(r, m[i])
	})
}

func Subtract[R ~[][]T, T types.Real](m R) types.MOperation[T] {
	return ForEachRow[T](func(i int, r []T) []T {
		return vector.Subtract(r, m[i])
	})
}

func Multiply[T types.Real](c T) types.MOperation[T] {
	return ForEachRow[T](func(i int, r []T) []T {
		return vector.Multiply(r, c)
	})
}

func Transpose[R ~[][]T, T types.Real](m R) types.MOperation[T] {
	return ForEachRow[T](func(i int, r []T) []T {
		pool.Put(r)
		return Column(m, i)
	})
}

//func Product[R ~[][]T, T types.Real](m1, m2 R) types.MOperation[T] {
//	if len(m1[0]) != len(m2) {
//		panic(errors.UnmatchedSizeOfMatricesError)
//	}
//
//	var transposed = Transpose(m2)
//
//	return ForEach[T](func(i, j int, t T) (r T) {
//		r = vector.DotProduct(m1[i], transposed[j])
//		return
//	})
//}

func Minor[R ~[][]T, T types.Real](values R, i, j int) types.MOperation[T] {
	var indexX = func(idx int) int {
		if idx < i {
			return idx
		}
		return idx + 1
	}
	return ForEachRow[T](func(i int, r []T) []T {
		copy(r[:j], values[indexX(i)][:j])
		copy(r[j:], values[indexX(i)][j+1:])
		return r
	})
}

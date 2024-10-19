package pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestGet", args{size: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual = Get[[]int](tt.args.size)
			assert.Equal(t, tt.args.size, len(actual))
		})
	}
}

func BenchmarkGet(b *testing.B) {
	var tests = []struct {
		Name string
		Size int
	}{
		{"Benchmark Get size 10", 10},
		{"Benchmark Get size 100", 100},
		{"Benchmark Get size 1000", 1000},
		{"Benchmark Get size 10000", 10000},
		{"Benchmark Get size 100000", 100000},
	}
	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				actual := Get[[]int](test.Size)
				Put(actual)
			}
		})
	}
}

func BenchmarkMake(b *testing.B) {
	var tests = []struct {
		Name string
		Size int
	}{
		{"Benchmark make size 10", 10},
		{"Benchmark make size 100", 100},
		{"Benchmark make size 1000", 1000},
		{"Benchmark make size 10000", 10000},
		{"Benchmark make size 100000", 100000},
	}
	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = make([]int, test.Size)
			}
		})
	}
}

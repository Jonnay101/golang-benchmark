package bench

import (
	"fmt"
	"reflect"
	"testing"
)

var fullnames = []string{
	"Paul Weller",
	"Jean Claude Van-Damme",
	"Robinho",
	"Theodore Smithy Grossen Wotsit",
	"",
}

var lastnames = []string{
	"Weller",
	"Van-Damme",
	"Robinho",
	"Wotsit",
}

var getLastNameFuncs = []testFunc{
	{name: "getLastnamesAppend", fn: getLastnamesAppend},
	{name: "getLastnamesIndex", fn: getLastnamesIndex},
}

type testFunc struct {
	name string
	fn   func([]string) []string
}

func Test_getLastnamesFuncs(t *testing.T) {
	for _, fn := range getLastNameFuncs {
		type args struct {
			fullnames []string
		}
		tests := []struct {
			name string
			args args
			want []string
		}{
			{fn.name, args{fullnames}, lastnames},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := fn.fn(tt.args.fullnames); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s() = %v, want %v", fn.name, got, tt.want)
				}
			})
		}
	}
}

func Benchmark_getLastNameFuncs(b *testing.B) {
	for _, fn := range getLastNameFuncs {
		b.Run(fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn.fn(fullnames)
			}
		})
	}
}

type addNumsFunc struct {
	name string
	fn   func(...int) int
}

var addNumsFuncs = []addNumsFunc{
	{"addNumsLoop", addNumsLoop},
	{"addNumsRecursive", addNumsRecursive},
}

func Test_addNumsFuncs(t *testing.T) {
	numList := createSliceOfInts(1, 1000, 1)
	for _, fn := range addNumsFuncs {
		type args struct {
			nums []int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			{fn.name, args{numList[:1]}, 1},
			{fn.name, args{numList[:8]}, 36},
			{fn.name, args{numList}, 499500},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := fn.fn(tt.args.nums...); got != tt.want {
					t.Errorf("%s() = %v, want %v", fn.name, got, tt.want)
				}
			})
		}
	}
}

func Benchmark_addNumsFuncs(b *testing.B) {
	nums := createSliceOfInts(0, 1000, 1)

	for _, fn := range addNumsFuncs {
		for idx := 0; idx <= 1000; idx += 100 {
			b.Run(fmt.Sprintf("%s-%d", fn.name, idx), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fn.fn(nums[:idx]...)
				}
			})
		}
	}
}

func createSliceOfInts(start, max, spacing int) []int {
	nums := make([]int, 0)
	for i := start; i < max; i += spacing {
		nums = append(nums, i)
	}
	return nums
}

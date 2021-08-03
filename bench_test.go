package bench

import (
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

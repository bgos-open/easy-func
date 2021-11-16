package is_array

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		v interface{}
	}
	type Ts struct {}
	var ts Ts

	tests := []struct {
		name string
		args args
		want bool
	}{
		// reference: https://github.com/php/php-src/blob/5b01c4863fe9e4bc2702b2bbf66d292d23001a18/ext/standard/tests/general_functions/is_array.phpt
		{
			name: "case-1:array-true",
			args: args{
				v: [1]string{""},
			},
			want: true,
		},
		{
			name: "case-2:array-true",
			args: args{
				v: [2]string{"string", "test"},
			},
			want: true,
		},
		{
			name: "case-3:array-true",
			args: args{
				v: [1]bool{true},
			},
			want: true,
		},
		{
			name: "case-4:array-true",
			args: args{
				v: [2]float32{10.5, 5.6},
			},
			want: true,
		},
		{
			name: "case-5:array-true",
			args: args{
				v: [][]int{},
			},
			want: true,
		},
		{
			name: "case-6:array-true",
			args: args{
				v: [2][3]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
				},
			},
			want: true,
		},
		{
			name: "case-7:slice-true",
			args: args{
				v: make([]int, 1),
			},
			want: true,
		},
		{
			name: "case-8:slice-true",
			args: args{
				v: [...]float32{},
			},
			want: true,
		},
		{
			name: "case-9:integers-false",
			args: args{
				v: -5322,
			},
			want: false,
		},
		{
			name: "case-10:integers-false",
			args: args{
				v: 0x55F,
			},
			want: false,
		},
		{
			name: "case-11:integers-false",
			args: args{
				v: -0xCCF,
			},
			want: false,
		},
		{
			name: "case-12:string-false",
			args: args{
				v: "",
			},
			want: false,
		},
		{
			name: "case-13:string-false",
			args: args{
				v: "string",
			},
			want: false,
		},
		{
			name: "case-14:string-false",
			args: args{
				v: ``,
			},
			want: false,
		},
		{
			name: "case-15:float-false",
			args: args{
				v: 10.0000000000000000005,
			},
			want: false,
		},
		{
			name: "case-16:float-false",
			args: args{
				v: .5e6,
			},
			want: false,
		},
		{
			name: "case-17:float-false",
			args: args{
				v: -.5E7,
			},
			want: false,
		},
		{
			name: "case-18:float-false",
			args: args{
				v: .5E+8,
			},
			want: false,
		},
		{
			name: "case-19:float-false",
			args: args{
				v: -.5e+90,
			},
			want: false,
		},
		{
			name: "case-20:float-false",
			args: args{
				v: 1e5,
			},
			want: false,
		},
		{
			name: "case-21:map-true",
			args: args{
				v: map[string]int{"nasa": 1},
			},
			want: true,
		},
		{
			name: "case-22:bool-false",
			args: args{
				v: true,
			},
			want: false,
		},
		{
			name: "case-23:bool-false",
			args: args{
				v: false,
			},
			want: false,
		},
		{
			name: "case-24:chan-false",
			args: args{
				v: make(chan string),
			},
			want: false,
		},
		{
			name: "case-25:struct-false",
			args: args{
				v: ts,
			},
			want: false,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
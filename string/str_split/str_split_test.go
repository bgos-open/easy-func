package str_split

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case 1:test string split",
			args: args{
				s: "test string split",
				n: 5,
			},
			want: []string{"test ","strin","g spl","it"},
		},
		{
			name: "case 2:empty string",
			args: args{
				s: "",
				n: 5,
			},
			want: []string{""},
		},
		{
			name: "case 3:chinese char",
			args: args{
				s: "中文字符串按长度分割为数组",
				n: 5,
			},
			want: []string{"中文字符串","按长度分割","为数组"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
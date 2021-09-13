package boolval

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// FROM EXAMPLE: https://github.com/php/php-src/blob/5b01c4863fe9e4bc2702b2bbf66d292d23001a18/ext/standard/tests/general_functions/boolval.phpt
		{
			name: "case-1:string-false",
			args: args{
				value: "",
			},
			want: false,
		},
		{
			name: "case-2:string-true",
			args: args{
				value: "-",
			},
			want: true,
		},
		{
			name: "case-3:int-false",
			args: args{
				value: 0,
			},
			want: false,
		},
		{
			name: "case-4:int-true",
			args: args{
				value: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.value); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

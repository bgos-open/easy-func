package ord

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/ord_basic.phpt
		{
			"case-1:a",
			args{`a`},
			97,
		},
		{
			"case-1.1:z",
			args{`z`},
			122,
		},
		{
			"case-2:0",
			args{`0`},
			48,
		},
		{
			"case-2.1:9",
			args{`9`},
			57,
		},
		{
			"case-3:!",
			args{`!`},
			33,
		},
		{
			"case-3.1:*",
			args{`*`},
			42,
		},
		{
			"case-3.2:@",
			args{`@`},
			64,
		},
		{
			"case-4:newline",
			args{"\n"},
			10,
		},
		{
			"case-5:hexadecimal escape character for newline",
			args{"\x0A"},
			10,
		},
		{
			"case-5.1:hexadecimal escape character boundary value",
			args{"\xFF"},
			255,
		},
		{
			"case-5.2:the maximum hexadecimal escape character of the regular ASCII code(Delete)",
			args{"\x7F"},
			127,
		},
		{
			"case-6:word",
			args{`Hello`},
			72,
		},
		// My use case
		{
			"case-7:H",
			args{`H`},
			72,
		},
		{
			"case-8:space",
			args{` `},
			32,
		},
		{
			"case-9:blank",
			args{``},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

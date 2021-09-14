package chr

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		ascii int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/chr_basic.phpt
		// var_dump('H' ===  chr(72));
		{
			"case-1:72",
			args{72},
			"H",
		},
		// var_dump('e' ===  chr(101));
		{
			"case-2:101",
			args{101},
			"e",
		},
		// var_dump('l' ===  chr(108));
		{
			"case-3:108",
			args{108},
			"l",
		},
		// var_dump('o' ===  chr(111));
		{
			"case-4:111",
			args{111},
			"o",
		},
		// var_dump('\n' ===  chr(10));
		{
			"case-5:newline",
			args{10},
			"\n",
		},
		// My use case
		{
			"case-1:number 0",
			args{0},
			"",
		},
		{
			"case-2:normal number",
			args{97},
			"a",
		},
		{
			"case-3.1:tab",
			args{9},
			"	",
		},
		{
			"case-4:delete",
			args{127},
			"\u007F",
		},
		{
			"case-5:over 255, remainder by 256",
			args{256},
			"",
		},
		{
			"case-5.1:negative number, add 256",
			args{-159},
			"a",
		},
		{
			"case-6:over 127",
			args{128},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.ascii); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

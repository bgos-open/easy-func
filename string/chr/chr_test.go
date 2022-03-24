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
		{
			"case-9:202",
			args{202},
			"",
		},
		{
			"case-10:29",
			args{29},
			"\u001d", // \u001d
		},
		{
			"case-11:54",
			args{54},
			"6",
		},
		{
			"case-12:160",
			args{160},
			"",
		},
		{
			"case-13:237",
			args{237},
			"",
		},
		{
			"case-14:27",
			args{27},
			"\u001b",
		},
		{
			"case-15:120",
			args{120},
			"x",
		},
		{
			"case-16:40",
			args{40},
			"(",
		},
		{
			"case-17:253",
			args{253},
			"",
		},
		{
			"case-18:149",
			args{149},
			"",
		},
		{
			"case-19:244",
			args{244},
			"",
		},
		{
			"case-20:0",
			args{0},
			"", // \u0000
		},
		{
			"case-21:99",
			args{99},
			"c",
		},
		{
			"case-22:235",
			args{235},
			"",
		},
		{
			"case-23:98",
			args{98},
			"b",
		},
		{
			"case-24:151",
			args{151},
			"",
		},
		{
			"case-25:129",
			args{129},
			"",
		},
		{
			"case-26:116",
			args{116},
			"t",
		},
		{
			"case-27:33",
			args{33},
			"!",
		},
		{
			"case-28:24",
			args{24},
			"\u0018", // \u0018
		},
		{
			"case-29:56",
			args{56},
			"8",
		},
		{
			"case-30:81",
			args{81},
			"Q",
		},
		{
			"case-31:85",
			args{85},
			"U",
		},
		{
			"case-32:7",
			args{7},
			"\u0007", //
		},
		{
			"case-33:206",
			args{206},
			"",
		},
		{
			"case-34:193",
			args{193},
			"",
		},
		{
			"case-35:65",
			args{65},
			"A",
		},
		{
			"case-36:229",
			args{229},
			"",
		},
		{
			"case-37:145",
			args{145},
			"",
		},
		{
			"case-38:33",
			args{33},
			"!",
		},
		{
			"case-39:171",
			args{171},
			"",
		},
		{
			"case-40:51",
			args{51},
			"3",
		},
		{
			"case-41:244",
			args{244},
			"",
		},
		{
			"case-42:195",
			args{195},
			"",
		},
		{
			"case-43:244",
			args{244},
			"",
		},
		{
			"case-44:12",
			args{12},
			"\f", //
		},
		{
			"case-45:163",
			args{163},
			"",
		},
		{
			"case-46:108",
			args{108},
			"l",
		},
		{
			"case-47:235",
			args{235},
			"",
		},
		{
			"case-48:111",
			args{111},
			"o",
		},
		{
			"case-49:184",
			args{184},
			"",
		},
		{
			"case-50:167",
			args{167},
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

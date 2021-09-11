package trim

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str       string
		character []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// 以下用例参考：https://github.com/php/php-src/blob/master/ext/standard/tests/strings/trim.phpt
		// var_dump('ABC' ===  trim('ABC'));
		{
			name: "case1:ABC",
			args: args{
				str: "ABC",
			},
			want: "ABC",
		},
		//var_dump('ABC' ===  trim(" \0\t\nABC \0\t\n"));
		{
			name: "case2:ABC",
			args: args{
				str: " \000\t\nABC \000\t\n",
			},
			want: "ABC",
		},
		//var_dump(" \0\t\nABC \0\t\n" ===  trim(" \0\t\nABC \0\t\n",''));
		{
			name: "case3:ABC",
			args: args{
				str:       " \000\t\nABC \000\t\n",
				character: []string{""},
			},
			want: " \000\t\nABC \000\t\n",
		},
		////var_dump("ABC\x50\xC1" === trim("ABC\x50\xC1\x60\x90","\x51..\xC0"));
		//{
		//	name: "case4:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x51..\xC0"},
		//	},
		//	want: "ABC\x50\xC1",
		//},
		////var_dump("ABC\x50" === trim("ABC\x50\xC1\x60\x90","\x51..\xC1"));
		//{
		//	name: "case5:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x51..\xC1"},
		//	},
		//	want: "ABC\x50",
		//},
		////var_dump("ABC" === trim("ABC\x50\xC1\x60\x90","\x50..\xC1"));
		//{
		//	name: "case6:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x50..\xC1"},
		//	},
		//	want: "ABC",
		//},
		////var_dump("ABC\x50\xC1" === trim("ABC\x50\xC1\x60\x90","\x51..\xC0"));
		//{
		//	name: "case7:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x50..\xC1"},
		//	},
		//	want: "ABC\x50\xC1",
		//},
		////var_dump("ABC\x50" === trim("ABC\x50\xC1\x60\x90","\x51..\xC1"));
		//{
		//	name: "case8:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x51..\xC1"},
		//	},
		//	want: "ABC\x50",
		//},
		////var_dump("ABC" === trim("ABC\x50\xC1\x60\x90","\x50..\xC1"));
		//{
		//	name: "case9:ABC",
		//	args: args{
		//		str:       "ABC\x50\xC1\x60\x90",
		//		character: []string{"\x50..\xC1"},
		//	},
		//	want: "ABC",
		//},
		// 以下是自己写的用例
		{
			name: "case-1:普通空格符",
			args: args{
				str:       " 12 3 ",
				character: []string{},
			},
			want: "12 3",
		},
		{
			name: "case-2.1:换行符等",
			args: args{
				str:       " 12 3 \n\r\t\v\0000\v\v\v",
				character: []string{},
			},
			want: "12 3",
		},
		{
			name: "case-2.2:换行符等",
			args: args{
				str:       " 12 3 \n\v",
				character: []string{},
			},
			want: "12 3",
		},
		{
			name: "case-3.1:不同的字符",
			args: args{
				str:       ` 12 3 abAc`,
				character: []string{"bacAcab"},
			},
			want: " 12 3 ",
		},
		{
			name: "case-4.1:空白能处理",
			args: args{
				str:       ``,
				character: []string{"bacAcab"},
			},
			want: "",
		},
		{
			name: "case-4.2:空白不去除",
			args: args{
				str:       ` `,
				character: []string{""},
			},
			want: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str, tt.args.character...); got != tt.want {
				t.Errorf("Trim() = |%s|, want |%s|", got, tt.want)
			}
		})
	}

}

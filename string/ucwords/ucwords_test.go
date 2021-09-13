package ucwords

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1:空格符分割",
			args: args{
				str: "testing ucwords",
			},
			want: "Testing Ucwords",
		},
		{
			name: "case-2:\t分割",
			args: args{
				str: "testing\tucwords",
			},
			want: "Testing	Ucwords",
		},
		{
			name: "case-3:空格符分割-2",
			args: args{
				str: " testing ucwords",
			},
			want: " Testing Ucwords",
		},
		{
			name: "case-4:\n分割",
			args: args{
				str: "testing\nucwords",
			},
			want: "Testing\nUcwords",
		},
		{
			name: "case-5:\v分割",
			args: args{
				str: "testing\vucwords",
			},
			want: "Testing\vUcwords",
		},
		{
			name: "case-6:两个空格符分割",
			args: args{
				str: "testing  ucwords",
			},
			want: "Testing  Ucwords",
		},
		{
			name: "case-7:\r分割",
			args: args{
				str: "testing\rucwords",
			},
			want: "Testing\rUcwords",
		},
		{
			name: "case-8:\f分割",
			args: args{
				str: "testing\fucwords",
			},
			want: "Testing\fUcwords",
		},
		{
			name: "case-9:多行字符串分割-1",
			args: args{
				str: ``,
			},
			want: "",
		},
		{
			name: "case-10:多行字符串分割-2",
			args: args{
				str: `

`,
			},
			want: `

`,
		},
		{
			name: "case-11:多行字符串分割-3",
			args: args{
				str: `
testing ucword() with
multiline string using
heredoc
`,
			},
			want: `
Testing Ucword() With
Multiline String Using
Heredoc
`,
		},
		// 注意：多行字符串不支持转义符
		{
			name: "case-12:多行字符串分割-4",
			args: args{
				str: "testing\rucword(str)\twith\n" +
					"multiline   string\t\tusing\n" +
					"heredoc\nstring.with\vdifferent\fwhite\vspaces",
			},
			want: "Testing\rUcword(Str)\tWith\n" +
				"Multiline   String\t\tUsing\n" +
				"Heredoc\nString.With\vDifferent\fWhite\vSpaces",
		},
		{
			name: "case-13:多行字符串分割-5",
			args: args{
				str: "12sting 123string 4567\n" +
					"multiline   string\t\tusing\n" +
					"heredoc\nstring.with\vdifferent\fwhite\vspaces",
			},
			want: "12sting 123string 4567\n" +
				"Multiline   String\t\tUsing\n" +
				"Heredoc\nString.With\vDifferent\fWhite\vSpaces",
		},
		{
			name: "case-14:多行字符串分割-6",
			args: args{
				str: "it's bright,but i cann't see it.\n" +
					"\"things in double quote\"\n" +
					"'things in single quote'\n" +
					"this\\line is /with\\slashs",
			},
			want: "It'S Bright,But I Cann'T See It.\n" +
				"\"Things In Double Quote\"\n" +
				"'Things In Single Quote'\n" +
				"This\\Line Is /With\\Slashs",
		},
		{
			name: "case-15:特殊符号",
			args: args{
				str: "!@#$%^&*()_+=-`~",
			},
			want: "!@#$%^&*()_+=-`~",
		},
		{
			name: "case-16:特殊符号-2",
			args: args{
				str: "t@@#$% %test ^test &test *test +test -test",
			},
			want: "T@@#$% %Test ^Test &Test *Test +Test -Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

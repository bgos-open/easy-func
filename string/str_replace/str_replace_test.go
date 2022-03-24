package strreplace

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		search  string
		replace string
		subject string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/str_replace_basic.phpt
		{
			name: `case-1: str_replace("", "", "")`,
			args: args{
				search:  "",
				replace: "",
				subject: "",
			},
			want: "",
		},
		{
			name: `case-2: str_replace("e", "b", "test")`,
			args: args{
				search:  "e",
				replace: "b",
				subject: "test",
			},
			want: "tbst",
		},
		{
			name: `case-3: str_replace("q", "q", "q")`,
			args: args{
				search:  "q",
				replace: "q",
				subject: "q",
			},
			want: "q",
		},
		{
			name: `case-4: str_replace("long string here", "", "")`,
			args: args{
				search:  "long string here",
				replace: "",
				subject: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.search, tt.args.replace, tt.args.subject); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

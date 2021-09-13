package trim

import "testing"

func TestCommon(t *testing.T) {
	type args struct {
		typeTrim TrimType
		s        string
		cutSets  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// FROM EXAMPLE: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/trim.phpt
		{
			name: "case-1:blank",
			args: args{
				typeTrim: DefaultTrim,
				s:        " 12 3 ",
				cutSets:  []string{},
			},
			want: "12 3",
		},
		{
			name: "case-2.1:newline",
			args: args{
				typeTrim: DefaultTrim,
				s:        " 12 3 \n\r\t\v\0000\v\v\v",
				cutSets:  []string{},
			},
			want: "12 3",
		},
		{
			name: "case-2.2:multi",
			args: args{
				typeTrim: DefaultTrim,
				s:        " 12 3 \n\v",
				cutSets:  []string{},
			},
			want: "12 3",
		},
		{
			name: "case-3.1:different char",
			args: args{
				typeTrim: DefaultTrim,
				s:        ` 12 3 abAc`,
				cutSets:  []string{"bacAcab"},
			},
			want: " 12 3 ",
		},
		{
			name: "case-4.1:none",
			args: args{
				typeTrim: DefaultTrim,
				s:        ``,
				cutSets:  []string{"bacAcab"},
			},
			want: "",
		},
		{
			name: "case-4.2:don't process",
			args: args{
				typeTrim: DefaultTrim,
				s:        ` `,
				cutSets:  []string{""},
			},
			want: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Common(tt.args.typeTrim, tt.args.s, tt.args.cutSets...); got != tt.want {
				t.Errorf("Trim() = |%s|, want |%s|", got, tt.want)
			}
		})
	}
}

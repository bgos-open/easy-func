//nolint
package substr

import "testing"

func TestSubstr(t *testing.T) {
	type args struct {
		str    string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/substr.phpt

		/* Variations for two arguments */
		{
			name: "case: empty,1",
			args: args{
				str:    "",
				start:  1,
				length: 6,
			},
			want: "",
		},
		{
			name: "case: empty,0",
			args: args{
				str:    "",
				start:  0,
				length: 6,
			},
			want: "",
		},
		{
			name: "case: empty,-2",
			args: args{
				str:    "",
				start:  -2,
				length: -6,
			},
			want: "",
		},
		{
			name: "case: abcdef,1",
			args: args{
				str:    "abcdef",
				start:  1,
				length: 6,
			},
			want: "bcdef",
		},
		{
			name: "case: abcdef,0",
			args: args{
				str:    "abcdef",
				start:  0,
				length: 6,
			},
			want: "abcdef",
		},
		{
			name: "case: abcdef,-2",
			args: args{
				str:    "abcdef",
				start:  -2,
				length: 6,
			},
			want: "ef",
		},

		/* Variations for three arguments */
		{
			name: "case: abcdef,1,3",
			args: args{
				str:    "abcdef",
				start:  1,
				length: 3,
			},
			want: "bcd",
		},
		{
			name: "case: abcdef,1,0",
			args: args{
				str:    "abcdef",
				start:  1,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: abcdef,1,-3",
			args: args{
				str:    "abcdef",
				start:  1,
				length: -3,
			},
			want: "bc",
		},
		{
			name: "case: abcdef,0,3",
			args: args{
				str:    "abcdef",
				start:  0,
				length: 3,
			},
			want: "abc",
		},
		{
			name: "case: abcdef,0,0",
			args: args{
				str:    "abcdef",
				start:  0,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: abcdef,0,-3",
			args: args{
				str:    "abcdef",
				start:  0,
				length: -3,
			},
			want: "abc",
		},
		{
			name: "case: abcdef,-2,3",
			args: args{
				str:    "abcdef",
				start:  -2,
				length: 3,
			},
			want: "ef",
		},
		{
			name: "case: abcdef,-2,0 ",
			args: args{
				str:    "abcdef",
				start:  -2,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: abcdef,-2,-3",
			args: args{
				str:    "abcdef",
				start:  -2,
				length: -3,
			},
			want: "",
		},
		{
			name: "case: 123abc,1,3",
			args: args{
				str:    "123abc",
				start:  1,
				length: 3,
			},
			want: "23a",
		},
		{
			name: "case: 123abc,1,0",
			args: args{
				str:    "123abc",
				start:  1,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: 123abc,1,-3",
			args: args{
				str:    "123abc",
				start:  1,
				length: -3,
			},
			want: "23",
		},
		{
			name: "case: 123abc,0,3",
			args: args{
				str:    "123abc",
				start:  0,
				length: 3,
			},
			want: "123",
		},
		{
			name: "case: 123abc,0,0",
			args: args{
				str:    "123abc",
				start:  0,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: 123abc,0,-3",
			args: args{
				str:    "123abc",
				start:  0,
				length: -3,
			},
			want: "123",
		},
		{
			name: "case: 123abc,-2,3",
			args: args{
				str:    "123abc",
				start:  -2,
				length: 3,
			},
			want: "bc",
		},
		{
			name: "case: 123abc,-2,0 ",
			args: args{
				str:    "123abc",
				start:  -2,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: 123abc,-2,-3",
			args: args{
				str:    "123abc",
				start:  -2,
				length: -3,
			},
			want: "",
		},
		{
			name: "case: _123abc,1,3",
			args: args{
				str:    "_123abc",
				start:  1,
				length: 3,
			},
			want: "123",
		},
		{
			name: "case: _123abc,1,0",
			args: args{
				str:    "_123abc",
				start:  1,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: _123abc,1,-3",
			args: args{
				str:    "_123abc",
				start:  1,
				length: -3,
			},
			want: "123",
		},
		{
			name: "case: _123abc,0,3",
			args: args{
				str:    "_123abc",
				start:  0,
				length: 3,
			},
			want: "_12",
		},
		{
			name: "case: _123abc,0,0",
			args: args{
				str:    "_123abc",
				start:  0,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: _123abc,0,-3",
			args: args{
				str:    "_123abc",
				start:  0,
				length: -3,
			},
			want: "_123",
		},
		{
			name: "case: _123abc,-2,3",
			args: args{
				str:    "_123abc",
				start:  -2,
				length: 3,
			},
			want: "bc",
		},
		{
			name: "case: _123abc,-2,0 ",
			args: args{
				str:    "_123abc",
				start:  -2,
				length: 0,
			},
			want: "",
		},
		{
			name: "case: _123abc,-2,-3",
			args: args{
				str:    "_123abc",
				start:  -2,
				length: -3,
			},
			want: "",
		},

		/* variation of start and length to point to same element */
		{
			name: "case: $str, 2, -2",
			args: args{
				str:    "abcde",
				start:  2,
				length: -2,
			},
			want: "c",
		},
		{
			name: "case: $str, -3, -2",
			args: args{
				str:    "abcde",
				start:  -3,
				length: -2,
			},
			want: "c",
		},

		/* Testing to return empty string when start denotes the position beyond the truncation (set by negative length) */
		{
			name: "case: $str, 4, -4",
			args: args{
				str:    "abcdef",
				start:  4,
				length: -4,
			},
			want: "",
		},

		/* Testing for string with null characters */
		{
			name: `case: abc\x0xy\x0z,2`,
			args: args{
				str:    `abc\x0xy\x0z`,
				start:  2,
				length: len(`abc\x0xy\x0z`),
			},
			want: `c\x0xy\x0z`,
		},

		/* start <0 && -start > length */
		{
			name: "case: abcd;',-8",
			args: args{
				str:    "abcd",
				start:  -8,
				length: 4,
			},
			want: "abcd",
		},

		// negative length
		{
			name: "case: negative length --- $str, 0, -1",
			args: args{
				str:    "abcdef",
				start:  0,
				length: -1,
			},
			want: "abcde",
		},
		{
			name: "case: $str, 1, -3",
			args: args{
				str:    "abcdef",
				start:  1,
				length: -3,
			},
			want: "bc",
		},
		{
			name: "case: negative length --- $str, 2, -1",
			args: args{
				str:    "abcdef",
				start:  2,
				length: -1,
			},
			want: "cde",
		},
		{
			name: "case: negative length --- $str, -3, -1",
			args: args{
				str:    "abcdef",
				start:  -3,
				length: -1,
			},
			want: "de",
		},

		// my case
		{
			name: "case: all string",
			args: args{
				str:    "abcde",
				start:  0,
				length: 5,
			},
			want: "abcde",
		},
		{
			name: "case: from 1 to 3",
			args: args{
				str:    "abcde",
				start:  1,
				length: 3,
			},
			want: "bcd",
		},
		{
			name: "case: over length",
			args: args{
				str:    "abcde",
				start:  2,
				length: 8,
			},
			want: "cde",
		},
		{
			name: "case: reverse from -1",
			args: args{
				str:    "abcde",
				start:  -1,
				length: 3,
			},
			want: "e",
		},
		{
			name: "case: reverse from -1",
			args: args{
				str:    "abcdef",
				start:  -1,
				length: 1,
			},
			want: "f",
		},
		{
			name: "case: reverse over length",
			args: args{
				str:    "abcde",
				start:  -2,
				length: 8,
			},
			want: "de",
		},
		{
			name: "case: include blank",
			args: args{
				str:    "ab cde",
				start:  1,
				length: 3,
			},
			want: "b c",
		},
		{
			name: "case: include chinese",
			args: args{
				str:    "ab好cde",
				start:  1,
				length: 3,
			},
			want: "b好c",
		},
		{
			name: "case: $str, -3, 1",
			args: args{
				str:    "abcdef",
				start:  -3,
				length: 1,
			},
			want: "d",
		},
		{
			name: "case: $str, 0, -10",
			args: args{
				str:    "abcdef",
				start:  -11,
				length: -10,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got1 := Substr(tt.args.str, tt.args.start, tt.args.length)
			// t.Logf("Substr() = %v, want %v", got1, tt.want)
			// got2 := Do(tt.args.str, tt.args.start, tt.args.length)
			// t.Logf("Do() = %v, want %v", got2, tt.want)
			// t.Log("================")
			if got := Substr(tt.args.str, tt.args.start, tt.args.length); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
			// if got := Do(tt.args.str, tt.args.start, tt.args.length); got != tt.want {
			// 	t.Errorf("Do() = %v, want %v", got, tt.want)
			// }
		})
	}
}

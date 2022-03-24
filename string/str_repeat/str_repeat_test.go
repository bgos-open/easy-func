package str_repeat

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		input      string
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/str_repeat.phpt
		{
			name: "case-1",
			args: args{
				input:      " ",
				multiplier: 1,
			},
			want: " ",
		},
		{
			name: "case-2",
			args: args{
				input:      " ",
				multiplier: 2,
			},
			want: "  ",
		},
		{
			name: "case-3",
			args: args{
				input:      " ",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-4",
			args: args{
				input:      " ",
				multiplier: -1,
			},
			want: "",
		},
		{
			name: "case-5",
			args: args{
				input:      "",
				multiplier: 1,
			},
			want: "",
		},
		{
			name: "case-6",
			args: args{
				input:      "",
				multiplier: 2,
			},
			want: "",
		},
		{
			name: "case-7",
			args: args{
				input:      "",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-8",
			args: args{
				input:      "",
				multiplier: -1,
			},
			want: "",
		},
		{
			name: "case-9",
			args: args{
				input:      "a",
				multiplier: 1,
			},
			want: "a",
		},
		{
			name: "case-10",
			args: args{
				input:      "a",
				multiplier: 2,
			},
			want: "aa",
		},
		{
			name: "case-11",
			args: args{
				input:      "a",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-12",
			args: args{
				input:      "a",
				multiplier: -1,
			},
			want: "",
		},
		{
			name: "case-13",
			args: args{
				input:      "%0",
				multiplier: 1,
			},
			want: "%0",
		},
		{
			name: "case-14",
			args: args{
				input:      "%0",
				multiplier: 2,
			},
			want: "%0%0",
		},
		{
			name: "case-15",
			args: args{
				input:      "%0",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-16",
			args: args{
				input:      "%0",
				multiplier: -1,
			},
			want: "",
		},
		{
			name: "case-17",
			args: args{
				input:      "1.23",
				multiplier: 1,
			},
			want: "1.23",
		},
		{
			name: "case-18",
			args: args{
				input:      "1.23",
				multiplier: 2,
			},
			want: "1.231.23",
		},
		{
			name: "case-19",
			args: args{
				input:      "1.23",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-20",
			args: args{
				input:      "1.23",
				multiplier: -1,
			},
			want: "",
		},
		{
			name: "case-21",
			args: args{
				input:      "\\0",
				multiplier: 1,
			},
			want: "\\0",
		},
		{
			name: "case-22",
			args: args{
				input:      "\\0",
				multiplier: 2,
			},
			want: "\\0\\0",
		},
		{
			name: "case-23",
			args: args{
				input:      "\\0",
				multiplier: 0,
			},
			want: "",
		},
		{
			name: "case-24",
			args: args{
				input:      "\\0",
				multiplier: -1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.input, tt.args.multiplier); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

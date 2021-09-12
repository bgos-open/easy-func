package str_pad

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		input     string
		padLength int
		padString string
		padType   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				input:     "str_pad()",
				padLength: 20,
				padString: "-+",
				padType:   strPadLeft,
			},
			want: "-+-+-+-+-+-str_pad()",
		},
		{
			name: "case-2",
			args: args{
				input:     "str_pad()",
				padLength: 20,
				padString: "-+",
				padType:   strPadRight,
			},
			want: "str_pad()-+-+-+-+-+-",
		},
		{
			name: "case-3",
			args: args{
				input:     "str_pad()",
				padLength: 20,
				padString: "-+",
				padType:   strPadBoth,
			},
			want: "-+-+-str_pad()-+-+-+",
		},
		{
			name: "case-4",
			args: args{
				input:     "variation",
				padLength: -1,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "variation",
		},
		{
			name: "case-5",
			args: args{
				input:     "variation",
				padLength: 0,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "variation",
		},
		{
			name: "case-6",
			args: args{
				input:     "variation",
				padLength: 9,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "variation",
		},
		{
			name: "case-7",
			args: args{
				input:     "variation",
				padLength: 10,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "=variation",
		},
		{
			name: "case-7",
			args: args{
				input:     "variation",
				padLength: 16,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "=======variation",
		},
		{
			name: "case-8",
			args: args{
				input:     "variation",
				padLength: -1,
				padString: "=",
				padType:   strPadRight,
			},
			want: "variation",
		},
		{
			name: "case-9",
			args: args{
				input:     "variation",
				padLength: 0,
				padString: "=",
				padType:   strPadRight,
			},
			want: "variation",
		},
		{
			name: "case-10",
			args: args{
				input:     "variation",
				padLength: 9,
				padString: "=",
				padType:   strPadRight,
			},
			want: "variation",
		},
		{
			name: "case-11",
			args: args{
				input:     "variation",
				padLength: 10,
				padString: "=",
				padType:   strPadRight,
			},
			want: "variation=",
		},
		{
			name: "case-12",
			args: args{
				input:     "variation",
				padLength: 16,
				padString: "=",
				padType:   strPadRight,
			},
			want: "variation=======",
		},
		{
			name: "case-13",
			args: args{
				input:     "variation",
				padLength: -1,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "variation",
		},
		{
			name: "case-14",
			args: args{
				input:     "variation",
				padLength: 0,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "variation",
		},
		{
			name: "case-15",
			args: args{
				input:     "variation",
				padLength: 9,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "variation",
		},
		{
			name: "case-16",
			args: args{
				input:     "variation",
				padLength: 10,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "variation=",
		},
		{
			name: "case-17",
			args: args{
				input:     "variation",
				padLength: 16,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "===variation====",
		},
		{
			name: "case-18",
			args: args{
				input:     "",
				padLength: -1,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "",
		},
		{
			name: "case-19",
			args: args{
				input:     "",
				padLength: 0,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "",
		},
		{
			name: "case-20",
			args: args{
				input:     "",
				padLength: 9,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "=========",
		},
		{
			name: "case-21",
			args: args{
				input:     "",
				padLength: 10,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "==========",
		},
		{
			name: "case-22",
			args: args{
				input:     "",
				padLength: 16,
				padString: "=",
				padType:   strPadLeft,
			},
			want: "================",
		},
		{
			name: "case-23",
			args: args{
				input:     "",
				padLength: -1,
				padString: "=",
				padType:   strPadRight,
			},
			want: "",
		},
		{
			name: "case-24",
			args: args{
				input:     "",
				padLength: 0,
				padString: "=",
				padType:   strPadRight,
			},
			want: "",
		},
		{
			name: "case-25",
			args: args{
				input:     "",
				padLength: 9,
				padString: "=",
				padType:   strPadRight,
			},
			want: "=========",
		},
		{
			name: "case-26",
			args: args{
				input:     "",
				padLength: 10,
				padString: "=",
				padType:   strPadRight,
			},
			want: "==========",
		},
		{
			name: "case-27",
			args: args{
				input:     "",
				padLength: 16,
				padString: "=",
				padType:   strPadRight,
			},
			want: "================",
		},
		{
			name: "case-28",
			args: args{
				input:     "",
				padLength: -1,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "",
		},
		{
			name: "case-29",
			args: args{
				input:     "",
				padLength: 0,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "",
		},
		{
			name: "case-30",
			args: args{
				input:     "",
				padLength: 9,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "=========",
		},
		{
			name: "case-31",
			args: args{
				input:     "",
				padLength: 10,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "==========",
		},
		{
			name: "case-32",
			args: args{
				input:     "",
				padLength: 16,
				padString: "=",
				padType:   strPadBoth,
			},
			want: "================",
		},
		{
			name: "case-33",
			args: args{
				input:     "variation",
				padLength: 16,
				padString: "=",
				padType:   3,
			},
			want: "variation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.input, tt.args.padLength, tt.args.padString, tt.args.padType); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

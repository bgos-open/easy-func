package empty

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// reference: https://github.com/php/php-src/blob/5b01c4863fe9e4bc2702b2bbf66d292d23001a18/ext/standard/tests/general_functions/boolval.phpt
		{
			name: "case-1",
			args: args{
				value: "",
			},
			want: true,
		},
		{
			name: "case-2",
			args: args{
				value: 0,
			},
			want: true,
		},
		{
			name: "case-3",
			args: args{
				value: 0.0,
			},
			want: true,
		},
		{
			name: "case-3",
			args: args{
				value: 0.00,
			},
			want: true,
		},
		{
			name: "case-3",
			args: args{
				value: 0.000,
			},
			want: true,
		},
		{
			name: "case-3",
			args: args{
				value: 0.0000,
			},
			want: true,
		},
		{
			name: "case-4",
			args: args{
				value: false,
			},
			want: true,
		},
		{
			name: "case-5",
			args: args{
				value: "test",
			},
			want: false,
		},
		{
			name: "case-5",
			args: args{
				value: " ",
			},
			want: false,
		},
		{
			name: "case-6",
			args: args{
				value: 1,
			},
			want: false,
		},
		{
			name: "case-7",
			args: args{
				value: 2,
			},
			want: false,
		},
		{
			name: "case-8",
			args: args{
				value: 0.01,
			},
			want: false,
		},
		{
			name: "case-9",
			args: args{
				value: 0.02,
			},
			want: false,
		},
		{
			name: "case-10",
			args: args{
				value: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.value); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

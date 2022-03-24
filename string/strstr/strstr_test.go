package strstr

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		haystack     string
		needle       string
		beforeNeedle bool
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
				haystack:     "test string",
				needle:       "test",
				beforeNeedle: false,
			},
			want: "test string",
		},
		{
			name: "case-2",
			args: args{
				haystack:     "test string",
				needle:       "string",
				beforeNeedle: false,
			},
			want: "string",
		},
		{
			name: "case-3",
			args: args{
				haystack:     "test string",
				needle:       "t s",
				beforeNeedle: false,
			},
			want: "t string",
		},
		{
			name: "case-4",
			args: args{
				haystack:     "tEst",
				needle:       "test",
				beforeNeedle: false,
			},
			want: "",
		},
		{
			name: "case-5",
			args: args{
				haystack:     "",
				needle:       "",
				beforeNeedle: false,
			},
			want: "",
		},
		{
			name: "case-6",
			args: args{
				haystack:     "a",
				needle:       "",
				beforeNeedle: false,
			},
			want: "",
		},
		{
			name: "case-7",
			args: args{
				haystack:     "",
				needle:       "a",
				beforeNeedle: false,
			},
			want: "",
		},
		{
			name: "case-8",
			args: args{
				haystack:     "aexample.com",
				needle:       "@",
				beforeNeedle: false,
			},
			want: "",
		},
		{
			name: "case-9",
			args: args{
				haystack:     "aexample.com",
				needle:       "@",
				beforeNeedle: true,
			},
			want: "",
		},
		{
			name: "case-10",
			args: args{
				haystack:     "a@example.com",
				needle:       "@",
				beforeNeedle: false,
			},
			want: "@example.com",
		},
		{
			name: "case-11",
			args: args{
				haystack:     "a@example.com",
				needle:       "@",
				beforeNeedle: true,
			},
			want: "a",
		},
		{
			name: "case-12",
			args: args{
				haystack:     "asdfasdfas@e",
				needle:       "@",
				beforeNeedle: false,
			},
			want: "@e",
		},
		{
			name: "case-13",
			args: args{
				haystack:     "asdfasdfas@e",
				needle:       "@",
				beforeNeedle: true,
			},
			want: "asdfasdfas",
		},
		{
			name: "case-14",
			args: args{
				haystack:     "@",
				needle:       "@",
				beforeNeedle: false,
			},
			want: "@",
		},
		{
			name: "case-15",
			args: args{
				haystack:     "@",
				needle:       "@",
				beforeNeedle: true,
			},
			want: "",
		},
		{
			name: "case-16",
			args: args{
				haystack:     "eE@fF",
				needle:       "e",
				beforeNeedle: false,
			},
			want: "eE@fF",
		},
		{
			name: "case-17",
			args: args{
				haystack:     "eE@fF",
				needle:       "e",
				beforeNeedle: true,
			},
			want: "",
		},
		{
			name: "case-18",
			args: args{
				haystack:     "eE@fF",
				needle:       "E",
				beforeNeedle: false,
			},
			want: "E@fF",
		},
		{
			name: "case-19",
			args: args{
				haystack:     "eE@fF",
				needle:       "E",
				beforeNeedle: true,
			},
			want: "e",
		},
		{
			name: "case-20",
			args: args{
				haystack:     "",
				needle:       " ",
				beforeNeedle: false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.haystack, tt.args.needle, tt.args.beforeNeedle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

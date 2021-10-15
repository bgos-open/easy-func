package strpos

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		haystack string
		needle   string
		offset   int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			"case-1:",
			args{
				"test string.",
				"test",
				0,
			},
			0,
		},
		{
			"case-2:",
			args{
				"test string.",
				"string",
				0,
			},
			5,
		},
		{
			"case-3:",
			args{
				"test string string",
				"string",
				8,
			},
			12,
		},
		{
			"case-4:",
			args{
				"test string",
				"hello",
				0,
			},
			false,
		},
		{
			"case-5:",
			args{
				"test string",
				"String",
				0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.haystack, tt.args.needle, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
package strpbrk

import (
	"reflect"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		//FROM EXAMPLE:https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strpbrk_basic.phpt
		// TODO: Add test cases.
		{
			"A",
			args{
				"This is a Simple text.",
				"mi",
			},
			"",
		},
		{
			"B",
			args{
				"This is a Simple text",
				"a",
			},
			"a Simple text",
		},
		{
			"C",
			args{
				"This is a Simple text",
				"A",
			},
			"",
		},
		{
			"D",
			args{
				"5",
				"5",
			},
			"5",
		},
		{
			"E",
			args{
				"This is a Simple text.",
				" ",
			},
			" is a Simple text.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Do(tt.args.haystack, tt.args.needle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

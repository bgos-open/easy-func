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
		// TODO: Add test cases.
		{
			"case-1:搜索字符在最前面",
			args{
				"test string.",
				"test",
			},
			"test string.",
		},
		{
			"case-2:搜索字符不在最前面",
			args{
				"test string.",
				"string",
			},
			"string.",
		},
		{
			"case-3:非从头开始搜索",
			args{
				"test string string",
				"string",
			},
			"string string",
		},
		{
			"case-4:搜索字符不存在",
			args{
				"test string",
				"hello",
			},
			false,
		},
		{
			"case-5:区分大小写",
			args{
				"test string",
				"String",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.haystack, tt.args.needle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

package str_shuffle

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case-1:test one word",
			args: args{
				str : "t",
			},
			want: "t",
		},
		{
			name: "case-2:test words",
			args: args{
				str : "Hello World！",
			},
			want: "Hello World！",
		},
		{
			name: "case-3:test Chinese and English",
			args: args{
				str : "你好abc",
			},
			want: "你好abc",
		},
		{
			name: "case-4:test empty word",
			args: args{
				str : "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got  := Do(tt.args.str)
			gotLen := len(got)
			if gotLen != len(tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

package str_contains

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case-1：normal",
			args: args{
				s: "test string",
				sub: "string",
			},
			want:true,
		},
		{
			name: "case-2：contain blank",
			args: args{
				s: "test string",
				sub: "t s",
			},
			want:true,
		},
		{
			name: "case-3：single char",
			args: args{
				s: "test string",
				sub: "g",
			},
			want:true,
		},
		{
			name: "case-4：tEst",
			args: args{
				s: "test string",
				sub: "tEst",
			},
			want:false,
		},
		{
			name: "case-5：empty string",
			args: args{
				s: "test string",
				sub: "",
			},
			want:true,
		},
		{
			name: "case-6：empty string",
			args: args{
				s: "",
				sub: "",
			},
			want:true,
		},
		{
			name: "case-7：empty string",
			args: args{
				s: "",
				sub: "t",
			},
			want:false,
		},
		{
			name: "case-8：escate char",
			args: args{
				s: "\\\\\\\\test string",
				sub: "\\test",
			},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
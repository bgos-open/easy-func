package hashmac

import (
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		input string
		key   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case-1",
			args: args{
				input: "GET?app_site_name=somyfo-com&com=tool&sessionid=bad4e102fa6469e42563df50b2fb7c83&ship_ab=1&t=testDemo",
				key:   "2ji2fl9s49a3c9f8891cb127243cf98a",
			},
			want: "5d5daac7930fe572afe3d2dfe0050c0790d2d98c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.input, tt.args.key); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

package join

import "testing"

func TestInt2string(t *testing.T) {
	type args struct {
		glue   string
		pieces []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/join_basic.phpt
		{
			name: "case: 1,2,3,4",
			args: args{
				glue:   ",",
				pieces: []int64{1, 2, 3, 4},
			},
			want: "1,2,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2string(tt.args.glue, tt.args.pieces); got != tt.want {
				t.Errorf("Int2string() = %v, want %v", got, tt.want)
			}
		})
	}
}

package is_bool

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
		{
			name: "case-1:",
			args: args{
				"aaa",
			},
			want: false,
		},
		{
			name: "case-2:",
			args: args{
				true,
			},
			want: true,
		},
		{
			name: "case-3:",
			args: args{
				false,
			},
			want: true,
		},
		{
			name: "case-4:",
			args: args{
				"aaa",
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

package isInt

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		values interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case-1",
			args: args{1},
			want: true,
		},
		{
			name: "case-2",
			args: args{"23"},
			want: false,
		},
		{
			name: "case-3",
			args: args{23.5},
			want: false,
		},
		{
			name: "case-4",
			args: args{true},
			want: false,
		},
		{
			name: "case-5",
			args: args{false},
			want: false,
		},
		{
			name: "case-6",
			args: args{uint(5)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.values); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

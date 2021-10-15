package isInt

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
			name: "case-1:string-false",
			args: args{
				value: "abc",
			},
			want: false,
		},
		{
			name: "case-2:int-true",
			args: args{
				value: 1,
			},
			want: true,
		},
		{
			name: "case-2:float-false",
			args: args{
				value: 1.01,
			},
			want: false,
		},
		{
			name: "case-2:slice-false",
			args: args{
				value: []int{4,5,6,7,8},
			},
			want: false,
		},
		{
			name: "case-2:map-false",
			args: args{
				value: map[string]string{"France": "Paris", "Italy": "Rome"},
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

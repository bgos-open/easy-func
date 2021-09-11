package boolval

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
		// TODO: Add test cases.
		{
			name: "case-1:string-false",
			args: args{
				value: "",
			},
			want: false,
		},
		{
			name: "case-2:string-true",
			args: args{
				value: "-",
			},
			want: true,
		},
		{
			name: "case-3:int-false",
			args: args{
				value: 0,
			},
			want: false,
		},
		{
			name: "case-4:int-true",
			args: args{
				value: 1,
			},
			want: true,
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

package isString

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
			name: "case string true",
			args: args{
				value: "string",
			},
			want: true,
		},
		{
			name: "case array false",
			args: args{
				value: []string{"1", "2"},
			},
			want: false,
		},
		{
			name: "case map false",
			args: args{
				value: map[string]string{
					"test1":"test1",
				},
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

package str_replace

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		old string
		new string
		s   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case 1:test string str_replace",
			args: args{
				old: "",
				new: "",
				s: "",
			},
			want: "",
		},
		{
			name: "case 2:test string str_replace",
			args: args{
				old: "e",
				new: "b",
				s: "test",
			},
			want: "tbst",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.old, tt.args.new, tt.args.s); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

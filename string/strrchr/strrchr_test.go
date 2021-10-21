package strrchr

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		s byte
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"a bc efg",
			args{'b'},
			"bc efg",
		},
		{
			"a bc efg",
			args{'e'},
			"efg",
		},
		{
			"a bc efg",
			args{' '},
			" efg",
		},
		{
			"a bc efg",
			args{'1'},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s, tt.name); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

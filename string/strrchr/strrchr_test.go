package strrchr

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		s interface{}
	}
	var k int8 = 97
	var k2 uint8 = 97
	var k3  = [2]int{1,2}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"a bc efg",
			args{"b"},
			"bc efg",
		},
		{
			"a bc efg",
			args{"bce"},
			"bc efg",
		},
		{
			"a bc efg",
			args{99},
			"c efg",
		},
		{
			"a bc efg",
			args{"e"},
			"efg",
		},
		{
			"a bc efg",
			args{k},
			"a bc efg",
		},
		{
			"a bc efg",
			args{k2},
			"a bc efg",
		},
		{
			"a bc efg",
			args{""},
			"",
		},
		{
			"a bc efg",
			args{k3},
			"",
		},
		{
			"a bc efg",
			args{12222},
			"",
		},
		{
			"a bc efg",
			args{" "},
			" efg",
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

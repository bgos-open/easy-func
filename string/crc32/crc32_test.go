package crc32

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		str string
	}
	var tests = []struct {
		name string
		args args
		want uint32
	}{
		{
			"case-1",
			args{"foo"},
			uint32(2356372769),
		},
		{
			"case-2",
			args{"bar"},
			uint32(1996459178),
		},
		{
			"case-3",
			args{"baz"},
			uint32(2015626392),
		},
		{
			"case-4",
			args{"grldsajkopallkjasd"},
			uint32(824412087),
		},
		{
			"case-空字符串",
			args{""},
			uint32(0),
		},
		{
			"case-空格",
			args{" "},
			uint32(3916222277),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.str); got != tt.want {
				t.Errorf("crc32('%s'):expect %d, actual %d", tt.args.str, tt.want, got)
			}
		})
	}
}

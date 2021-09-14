package chunk_split

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		body     string
		chunklen uint
		end      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/chunk_split.phpt
		{
			"case-1:single character separation",
			args{`abc`, 1, `-`},
			"a-b-c-",
		},
		{
			"case-2:default end separation",
			args{`foooooooooooooooo`, 5, ``},
			"foooo\r\nooooo\r\nooooo\r\noo\r\n",
		},
		{
			"case-3:default split size",
			args{`XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`, 0, ``},
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\r\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\r\n",
		},
		{
			"case-4:The length of the split string is less than the split size",
			args{`test`, 10, `|end`},
			"test|end",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.body, tt.args.chunklen, tt.args.end); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

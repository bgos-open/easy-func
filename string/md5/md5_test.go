package md5

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Reference: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5.phpt
		{
			"case-1:blank",
			args{``},
			"d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			"case-1.1:space",
			args{` `},
			"7215ee9c7d9dc229d2921a40e899ec5f",
		},
		{
			"case-2:single character",
			args{`a`},
			"0cc175b9c0f1b6a831c399e269772661",
		},
		{
			"case-2.1:multiple characters",
			args{`apple`},
			"1f3870be274f6c49b3e31a0c6728957f",
		},
		{
			"case-3:string with spaces",
			args{`message digest`},
			"f96b697d7cb7938d525a2f31aaf161d0",
		},
		{
			"case-4:all lowercase string",
			args{`abcdefghijklmnopqrstuvwxyz`},
			"c3fcd3d76192e4007dfb496cca67e13b",
		},
		{
			"case-5:contains uppercase and lowercase strings",
			args{`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`},
			"d174ab98d277d9f5a5611c2c9f419d9f",
		},
		{
			"case-6:numeric string",
			args{`12345678901234567890123456789012345678901234567890123456789012345678901234567890`},
			"57edf4a22be3c955ac49da2e2107b67a",
		},
		// My use case
		{
			"case-7:contains a newline at the end(windows system)",
			args{"apple\r\n"},
			"979a0e192a27373e913c29a7b2477dae",
		},
		{
			"case-7.1:contains a newline at the end(linux/unix system)",
			args{"apple\n"},
			"30c6677b833454ad2df762d3c98d2409",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do(tt.args.s); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

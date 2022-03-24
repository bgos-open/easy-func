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
		{
			"case-8.1:random",
			args{"dsaff"},
			"364e569fdc48b433c95dda58009d3737",
		},
		{
			"case-8.2:random2",
			args{"1637604408"},
			"cad08d50195035a936a1150edb90b99b",
		},
		{
			"case-8.3:random3",
			args{"111"},
			"698d51a19d8a121ce581499d7b701668",
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

/*
Comparison of benchmark test results of three type methods
goos: windows
goarch: amd64
pkg: easy-func/string/md5
cpu: Intel(R) Core(TM) i5-6500 CPU @ 3.20GHz
BenchmarkDo
BenchmarkDo-4    	 2105058	       567.5 ns/op
BenchmarkDo2
BenchmarkDo2-4   	  876865	      1391 ns/op
BenchmarkDo3
BenchmarkDo3-4   	 1328520	       888.6 ns/op
PASS
*/

var str = "yh02oO0beCj9lf1UYe9mxdN0dhtqAVUBnMpB54orviv6zaM"

func BenchmarkDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Do(str)
	}
}
func BenchmarkDo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Do2(str)
	}
}
func BenchmarkDo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Do3(str)
	}
}

func TestDo2(t *testing.T) {
	type args struct {
		str string
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
		{
			"case-8.1:random",
			args{"dsaff"},
			"364e569fdc48b433c95dda58009d3737",
		},
		{
			"case-8.2:random2",
			args{"1637604408"},
			"cad08d50195035a936a1150edb90b99b",
		},
		{
			"case-8.3:random3",
			args{"111"},
			"698d51a19d8a121ce581499d7b701668",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do2(tt.args.str); got != tt.want {
				t.Errorf("Do2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo3(t *testing.T) {
	type args struct {
		str string
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
		{
			"case-8.1:random",
			args{"dsaff"},
			"364e569fdc48b433c95dda58009d3737",
		},
		{
			"case-8.2:random2",
			args{"1637604408"},
			"cad08d50195035a936a1150edb90b99b",
		},
		{
			"case-8.3:random3",
			args{"111"},
			"698d51a19d8a121ce581499d7b701668",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Do3(tt.args.str); got != tt.want {
				t.Errorf("Do3() = %v, want %v", got, tt.want)
			}
		})
	}
}

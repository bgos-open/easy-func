package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// Do https://www.php.net/manual/en/function.md5
func Do(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
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

func Do2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Do3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

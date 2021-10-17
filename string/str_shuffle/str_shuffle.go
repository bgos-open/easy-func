package str_shuffle

import (
	"math/rand"
	"time"
)
// https://www.php.net/manual/en/function.str-shuffle.php
func Do(str string) string {
	strRunes := []rune(str)
	strLen := len(strRunes)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s:= make([]rune, strLen)

	for i,v := range r.Perm(strLen) {
		s[i] = strRunes[v]
	}
	return  string(s)
}

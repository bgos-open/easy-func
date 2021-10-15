package strcspn

import (
	"math"
	"strings"
)

//https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strcspn.phpt
func Do(str string, char string, offset int, length int) int {
	// abs offset must less than the length of str
	if int(math.Abs(float64(offset))) >= len(str) {
		return 0
	}

	if offset > 0 {
		str = str[offset:]
	} else if offset < 0 {
		str = str[len(str) + offset:]
	}

	if length > 0 && length < len(str) {
		str = str[0: length]
	}

	for k,v := range str {
		index := strings.Index(char, string(v))
		if index != -1 {
			return k
		}
	}

	return  len(str)
}

package strpos

import "strings"

//https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strlen.phpt
func Do(haystack string, needle string, offset int) interface{} {
	if (offset > 0) {
		haystack = haystack[offset:]
	}
	r := strings.Index(haystack, needle)
	if r == -1 {
		return false
	}
	return  r + offset
}

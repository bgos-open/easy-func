package strcspn

import (
	"strings"
)

//https://github.com/php/php-src/blob/master/ext/standard/tests/strings/strcspn.phpt
func Do(str string, char string, offset int, length int) int {
	lenStr := len(str)
	if offset >= lenStr {
		return 0
	}
	if offset < 0 {
		offset = lenStr + offset
	}
	if length < 0 {
		length = lenStr + length - offset
	}
	if length <= 0 {
		return 0
	}
	reach := offset + length
	if reach >= lenStr {
		reach = lenStr
	}
	str = str[offset: reach]
	for k,v := range str {
		index := strings.Index(char, string(v))
		if index != -1 {
			return k
		}
	}
	return  len(str)
}

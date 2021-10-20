package strncmp

import (
	"errors"
	"strings"
)

func Do(str1, str2 string, length int) (res int, err error) {
	if length < 0 {
		return 0, errors.New("length must be greater than or equal to 0")
	} else if length == 0 {
		return 0, nil
	} else {
		len1, len2 := len(str1), len(str2)
		var minLen int
		if len1 <= len2 {
			minLen = len1
		} else {
			minLen = len2
		}
		if length > minLen {
			length = minLen
		}
		return strings.Compare(str1[:length], str2[:length]), nil
	}

}

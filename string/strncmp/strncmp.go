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
		if len1 < len2 {
			length = len1
		} else {
			length = len2
		}
		return strings.Compare(str1[:length], str2[:length]), nil
	}

}

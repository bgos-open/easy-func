package strpbrk

import "strings"

func Do(haystack string, charList string) interface{} {
	if len(charList) == 0 {
		return false
	}
	r := strings.Index(haystack, charList)
	if r == -1 {
		return false
	}
	return haystack[r:]
}

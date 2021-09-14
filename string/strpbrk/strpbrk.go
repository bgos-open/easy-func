package strpbrk

import "strings"

func Do(haystack ...string) interface{} {
	if len(haystack) < 2 {
		return false
	}
	r := strings.Index(haystack[0], haystack[1])
	if r == -1 {
		return false
	}
	return haystack[r:]
}

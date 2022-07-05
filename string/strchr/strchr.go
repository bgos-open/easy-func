package strchr

import (
	"strings"
)

// https://www.php.net/manual/en/function.strchr
func Do(haystack string, needle string, beforeNeedle bool) string {
	/**
		In PHP, haystack or needle returns false when entering an empty string
	 */
	if haystack == "" || needle == "" {
		return ""
	}
	/**
		If needle is an empty string, the value of "idx" will be 0
	 */
	idx := strings.Index(haystack, needle)

	if idx == -1 {
		return ""
	}
	if beforeNeedle {
		return haystack[:idx]
	}
	return haystack[idx:]
}
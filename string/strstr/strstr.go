package strstr

import (
	"strings"
)

// https://www.php.net/manual/en/function.strstr
func Do(haystack string, needle string, beforeNeedle bool) string {
	if needle == "" {
		return ""
	}

	idx := strings.Index(haystack, needle)

	if idx == -1 {
		return ""
	}
	if beforeNeedle {
		return haystack[:idx]
	}
	return haystack[idx:]
}

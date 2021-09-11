package rtrim

import (
	"strings"
)

// https://www.php.net/manual/en/function.trim
func Do(s string, cutSets ...string) string {
	if len(cutSets) > 0 {
		return strings.TrimRight(s, cutSets[0])
	}

	return strings.TrimRight(s, " \r\n\t\v\u00000")
}

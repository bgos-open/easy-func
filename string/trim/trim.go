package trim

import (
	"strings"
)

// https://www.php.net/manual/en/function.trim
func Do(s string, cutSets ...string) string {

	if len(cutSets) > 0 {
		return strings.Trim(s, cutSets[0])
	}

	return strings.Trim(s, " \r\n\t\v\u00000")
}

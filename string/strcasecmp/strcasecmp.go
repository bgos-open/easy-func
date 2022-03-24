package strcasecmp

import (
	"strings"
)

// https://www.php.net/manual/en/function.strcasecmp
func Do(string1 string, string2 string) int {
	if strings.EqualFold(string1, string2) {
		return 0
	}
	return 1
}

package str_repeat

import (
	"strings"
)

// https://www.php.net/manual/en/function.str-repeat
func Do(input string, multiplier int) string {
	if multiplier <= 0 {
		return ""
	}
	return strings.Repeat(input, multiplier)
}

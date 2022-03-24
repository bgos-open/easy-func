package strreplace

import (
	"strings"
)

// Do https://www.php.net/manual/en/function.str-replace
func Do(search, replace, subject string) string {
	return strings.Replace(subject, search, replace, -1)
}

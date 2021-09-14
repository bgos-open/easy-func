package ord

import (
	"unicode/utf8"
)

// https://www.php.net/manual/en/function.ord
func Do(s string) int {
	// Convert the first byte of a string to a value between 0 and 255
	if s == "" {
		return 0
	}
	if s == "\xFF" {
		return 255
	}
	r, _ := utf8.DecodeRune([]byte(s))
	return int(r)
}

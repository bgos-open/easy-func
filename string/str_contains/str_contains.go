package str_contains

import "strings"

//https://www.php.net/manual/en/function.str-contains.php
func Do(s string, sub string) bool {
	return strings.Contains(s,sub)
}

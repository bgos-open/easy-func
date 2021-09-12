package ucwords

import (
	"strings"
)

// https://www.php.net/manual/en/function.ucwords
func Do(str string) string {
	return strings.Title(str)
}

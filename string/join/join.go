package join

import (
	"strings"
)

// https://www.php.net/manual/zh/function.join.php
func Do(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

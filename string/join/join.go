package join

import (
	"strconv"
	"strings"
)

// Int2string https://www.php.net/manual/zh/function.join
func Int2string(glue string, pieces []int64) string {
	var IDs []string
	for _, i := range pieces {
		IDs = append(IDs, strconv.FormatInt(i, 10))
	}
	return strings.Join(IDs, glue)
}

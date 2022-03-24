package strrchr

import (
	"strings"
)

func Do(s byte, sub string) string {
	subStr := string(s)
	index := strings.LastIndex(sub, subStr)
	if index != -1 {
		return sub[index:]
	}
	return ""
}

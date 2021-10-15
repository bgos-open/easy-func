package str_replace

import (
	"strings"
)

func Do(old string,new string,s string) string {

	return strings.Replace(s, old, new,-1)
}
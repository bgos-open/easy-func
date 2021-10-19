package implode

import (
	"strings"
)

// Do — implode — Join array elements with a string (https://www.php.net/manual/en/function.implode.php)
// Tip: The merging of the map type is affected by the return order of reflect.keys, and the result of the merging will be different from the order of the map.
func Do(Sep string, array []string) (res string) {
	strLen := len(array)
	if strLen == 0 {
		return  ""
	}
	if strLen == 1 {
		return array[0]
	}
	var buff strings.Builder
	buff.WriteString(array[0])
	for _, s := range array[1:] {
		buff.WriteString(Sep)
		buff.WriteString(s)
	}
	return buff.String()

}

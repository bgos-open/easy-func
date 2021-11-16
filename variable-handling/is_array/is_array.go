package is_array

import (
	"reflect"
)

func Do(v interface{}) bool {
	vType := reflect.TypeOf(v).Kind().String()
	return vType == "array" || vType == "slice" || vType == "map"
}
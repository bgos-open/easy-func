package is_array

import (
	"reflect"
)

func Do(v interface{}) bool {
	vType := reflect.TypeOf(v).Kind()
	return vType == reflect.Array || vType == reflect.Slice || vType == reflect.Map

}
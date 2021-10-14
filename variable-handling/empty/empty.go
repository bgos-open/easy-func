package empty

import (
	"reflect"
)

// https://www.php.net/manual/zh/function.empty.php
func Do(value interface{}) bool {
	if reflect.TypeOf(value).String() == "string" && value == "" {
		return true
	}

	if reflect.TypeOf(value).String() == "int" && value == 0 {
		return true
	}

	if reflect.TypeOf(value).String() == "float32" && value == 0.0 {
		return true
	}

	if reflect.TypeOf(value).String() == "float64" && value == 0.0 {
		return true
	}

	if reflect.TypeOf(value).String() == "bool" && value == false {
		return true
	}

	return false
}

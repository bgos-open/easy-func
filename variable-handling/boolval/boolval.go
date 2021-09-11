package boolval

import (
	"reflect"
)

// https://www.php.net/manual/zh/function.boolval.php
func Do(value interface{}) bool {
	if reflect.TypeOf(value).String() == "string" && value == "" {
		return false
	}

	if reflect.TypeOf(value).String() == "int" && value == 0 {
		return false
	}

	return true
}

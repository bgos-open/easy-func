package isInt

import "reflect"

// https://www.php.net/manual/en/function.is-int
func Do(values interface{}) bool {
	of := reflect.TypeOf(values)
	if of.Kind() == reflect.Int || of.Kind() == reflect.Int8  || of.Kind() == reflect.Int16  || of.Kind() == reflect.Int32  || of.Kind() == reflect.Int64 {
		return true
	}

	if of.Kind() == reflect.Uint || of.Kind() == reflect.Uint8  || of.Kind() == reflect.Uint16 || of.Kind() == reflect.Uint32  || of.Kind() == reflect.Uint64 {
		return true
	}
	return  false

}

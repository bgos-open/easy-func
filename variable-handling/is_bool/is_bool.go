package is_bool
import (
	"reflect"
)

func Do(value interface{}) bool{
	if reflect.TypeOf(value).String() == "bool"{
		return true
	}
	return false
}

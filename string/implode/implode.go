package implode

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Do — implode — Join array elements with a string (https://www.php.net/manual/en/function.implode.php)
// Tip: The merging of the map type is affected by the return order of reflect.keys, and the result of the merging will be different from the order of the map.
func Do(Sep string, array interface{}) (res string, err error) {
	of := reflect.TypeOf(array)
	valueOf := reflect.ValueOf(array)
	var buff strings.Builder
	switch of.Kind() {
		case reflect.Array,reflect.Slice:
			len := valueOf.Len()
			for i:= 0 ; i<= len - 1; i++ {
				toString := dataToString(valueOf.Index(i).Interface())
				buff.WriteString(toString)
				if i < len - 1 && len > 1  {
					buff.WriteString(Sep)
				}
			}
			res = buff.String()
	case reflect.Map:
		mapKey := valueOf.MapKeys()
		mapLen := len(mapKey)
		for i := 0 ; i < mapLen; i ++ {
			toString := dataToString(valueOf.MapIndex(mapKey[i]).Interface())
			buff.WriteString(toString)
			if i < mapLen - 1 && mapLen > 1  {
				buff.WriteString(Sep)
			}
		}
		res = buff.String()
	default:
		err = errors.New("invalid arguments array")
		break
	}
	return
}

func dataToString(str interface{})(resStr string){
	kind := reflect.TypeOf(str).Kind()
	valueOf := reflect.ValueOf(str)
	if kind == reflect.String {
		resStr = valueOf.String()
	}else if kind == reflect.Int || kind == reflect.Int32  || kind == reflect.Int16  || kind == reflect.Int64 {
		resStr = strconv.FormatInt(valueOf.Int(),10)
	}else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint32  || kind == reflect.Uint16  || kind == reflect.Uint64 {
		resStr = strconv.FormatUint(valueOf.Uint(),10)
	}else if kind == reflect.Float32 || kind == reflect.Float64{
		resStr = strconv.FormatFloat(valueOf.Float(), 'f', -1,64)
	}else{
		resStr = reflect.TypeOf(str).Kind().String()
	}
	return
}
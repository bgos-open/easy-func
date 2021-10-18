package strrchr

import (
	"reflect"
	"strconv"
	"strings"
)

func getSting(s interface{})string{
	kind := reflect.TypeOf(s).Kind()
	valueOf := reflect.ValueOf(s)
	var s2 string
	if kind == reflect.String {
		s2 = valueOf.String()
	}else if kind == reflect.Int || kind == reflect.Int32  || kind == reflect.Int16  || kind == reflect.Int64  || kind == reflect.Int8 {
		changeStr := strconv.FormatInt(valueOf.Int(),10)
		sInt, _ := strconv.Atoi(changeStr)
		s2 = string(rune(sInt))
	}else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint32  || kind == reflect.Uint16  || kind == reflect.Uint64 {
		changeStr := strconv.FormatUint(valueOf.Uint(),10)
		sInt, _ := strconv.Atoi(changeStr)
		s2 = string(rune(sInt))
	}

	return s2
}

func Do(s interface{},sub string) string {

	s3 := getSting(s)
	if len(s3) == 0{
		return ""
	}

	str := s3[:1]
	index := strings.LastIndex(sub,str)
	if index != -1{
		return sub[index:]
	}
	return ""
}


package stristr

import (
	"errors"
	"fmt"
	"strings"
)
//Do https://github.com/php/php-src/blob/master/ext/standard/tests/strings/stristr.phpt
func Do(str string,sep string, before_search bool) (string, error) {
	if sep == "" {
		return  "", errors.New("empty needle param sep")
	}
	index := strings.Index(strings.ToLower(str), strings.ToLower(sep))
	if index == -1 {
		return  "", errors.New(fmt.Sprintf("param %s is not exits", sep))
	}
	strArray := make([]string, 0)
	if before_search {
		strArray = strings.Split(str, "")[:index]
	}else{
		strArray = strings.Split(str, "")[index:]

	}
	return strings.Join(strArray, ""), nil


}

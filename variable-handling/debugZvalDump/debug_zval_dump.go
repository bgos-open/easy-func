package debugZvalDump

import (
	"fmt"
)

/**
Dumps a string representation of an internal zval (Zend value) structure to output. This is mostly useful for
understanding or debugging implementation details of the Zend Engine or PHP extensions.
Since golang is not Zend, an alternative is used
*/

// https://www.php.net/manual/zh/function.debug-zval-dump.php
func Do(values ...interface{}) {
	for i, _ := range values {
		fmt.Printf("%#v\n", values[i])
	}
}

package nl2br

import (
	"bytes"
)

// https://www.php.net/manual/en/function.nl2br
func Do(str string, isXhtml bool) string {
	r,n := '\r','\n'
	strRunes := []rune(str)
	var strBr []byte
	var buf bytes.Buffer

	if !isXhtml {
		strBr = []byte("<br>")
	} else {
		strBr = []byte("<br />")
	}
	skip := false
	strLength := len(strRunes)
	for i,v := range strRunes {
		if skip {
			skip = false
			continue
		}

		switch v {
		case n,r:
			buf.Write(strBr)
			buf.WriteRune(v)
			if (i+1 < strLength) && ( (v == r && strRunes[i+1] == n) || (v == n && strRunes[i+1] == r) ) {
				buf.WriteRune(strRunes[i+1])
				skip = true
			}
		default:
			buf.WriteRune(v)
		}
	}
	return buf.String()
}
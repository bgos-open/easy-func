package isString

func Do(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	default :
		return false
	}
}
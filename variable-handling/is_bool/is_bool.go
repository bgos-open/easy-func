package is_bool

func Do(value interface{}) bool{
	switch value.(type) {
	case bool:
		return true
	default:
		return false
	}
}

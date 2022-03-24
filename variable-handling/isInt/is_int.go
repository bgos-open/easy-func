package isInt

//https://www.php.net/manual/zh/function.is-int.php
func Do(value interface{}) bool {
	switch value.(type) {
	case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint:
		return true
	default:
		return false
	}
}

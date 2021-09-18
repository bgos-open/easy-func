package strrev

func Do(str string) string {
	strLen := len(str)
	if len(str) <= 1 {
		return str
	}
	bStr := []byte(str)
	left := 0           //
	right := strLen - 1 //
	for i := 0; i < (strLen / 2); i++ {
		if left >= right {
			break
		}
		bStr[left], bStr[right] = bStr[right], bStr[left]
		left++
		right--
	}
	return string(bStr)
}

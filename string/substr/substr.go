package substr

// Substr 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start, length int) string {
	if start < 0 && length < 0 && start > length {
		return ""
	}

	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl + start
	}
	// 长度为负数
	if length < 0 {
		end = rl + length
	} else {
		end = start + length
		if end < length {
			end = length
		}
	}
	start, end = checkValue(start, rl, end)

	if start > end {
		return ""
	}
	return string(rs[start:end])
}

func checkValue(start, rl, end int) (newStart, newEnd int) {
	newStart, newEnd = start, end
	if start < 0 {
		newStart = 0
	}
	if start > rl {
		newStart = rl
	}
	if end < 0 {
		newEnd = 0
	}
	if end > rl {
		newEnd = rl
	}
	return newStart, newEnd
}

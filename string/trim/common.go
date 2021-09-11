package trim

type TrimType uint8

const (
	DefaultTrim TrimType = iota
	LeftTrim
	RightTrim
)

// https://www.php.net/manual/en/function.trim
func Common(typeTrim TrimType, s string, cutSets ...string) string {

	if len(cutSets) == 0 {
		cutSets = append(cutSets, " \n\r\t\v\0000")
	}

	var cutSetMap = map[string]struct{}{}
	var lengthCutSet = len(cutSets[0])
	for i := 0; i < lengthCutSet; i++ {
		cutSetMap[cutSets[0][i:i+1]] = struct{}{}
	}

	var length = len(s)
	leftTrimPos := 0
	if typeTrim != RightTrim {
		for ; leftTrimPos < length; leftTrimPos++ {
			if _, ok := cutSetMap[s[leftTrimPos:leftTrimPos+1]]; !ok {
				break
			}
		}
	}

	rightTrimPos := length - 1
	if typeTrim != LeftTrim && leftTrimPos != rightTrimPos {
		for ; rightTrimPos >= 0; rightTrimPos-- {
			if _, ok := cutSetMap[s[rightTrimPos:rightTrimPos+1]]; !ok {
				break
			}
		}
	}

	return s[leftTrimPos : rightTrimPos+1]
}

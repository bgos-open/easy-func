package str_pad

const strPadLeft = 0
const strPadRight = 1
const strPadBoth = 2

// https://www.php.net/manual/en/function.str-pad
func Do(input string, padLength int, padString string, padType int) string {
	inputLength := len(input)
	if padLength <= 0 || padLength <= inputLength {
		return input
	}

	numPadChars := padLength - inputLength

	leftPad, rightPad := 0, 0
	switch padType {
	case strPadRight:
		rightPad = numPadChars
		break
	case strPadLeft:
		leftPad = numPadChars
		break
	case strPadBoth:
		leftPad = numPadChars / 2
		rightPad = numPadChars - leftPad
		break
	default:
		return input
	}
	padStringLength := len(padString)

	leftString := ""
	rightString := ""
	for i := 0; i < leftPad; i++ {
		leftString += string(padString[i%padStringLength])
	}

	input = leftString + input

	for i := 0; i < rightPad; i++ {
		rightString += string(padString[i%padStringLength])
	}

	input += rightString

	return input
}

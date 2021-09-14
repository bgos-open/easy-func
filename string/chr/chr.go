package chr

// https://www.php.net/manual/en/function.chr
func Do(ascii int) string {
	// return string(ascii)
	/*
		https://www.php.net/manual/zh/function.chr.php#91868
		chr(-number) is equal to chr((number%256)+256)

		https://www.php.net/manual/zh/function.chr.php#41471
		Note that if the number is higher than 256, it will return the number mod 256.
		For example :
		chr(321)=A because A=65(256)
	*/
	if ascii < 0 {
		ascii += 256
	}
	ascii = ascii % 256

	// chr(0)===>NULL char
	if ascii == 0 {
		return ""
	}

	/*
		The results in PHP and Go are different because, as the documentation for each states,
		PHP's chr returns the ASCII character for its argument, whereas Go's rune uses UTF-8.
		Below 127, ASCII and UTF-8 are the same, but above that, they differ.
	*/
	// 0..127
	if ascii > 127 {
		return ""
	}
	return string(rune(ascii)) // 返回指定的字符
}

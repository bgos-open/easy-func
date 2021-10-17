package str_split

//https://www.php.net/manual/en/function.str-split.php
func Do(s string, n int) []string{
	var split []string

	runes := []rune(s)
	l:= len(runes)
	if l == 0 {
		return []string{s}
	}
	for i := 0; i < l; i += n {
		nn := i + n
		if nn > l {
			nn = l
		}
		split = append(split, string(runes[i:nn]))
	}
	return split
}

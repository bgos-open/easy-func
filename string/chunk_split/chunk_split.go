package chunk_split

// Split a string into smaller chunks
// https://www.php.net/manual/en/function.chunk-split
func Do(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	// Which is set to 76 characters long chunk and "\r\n" at the end of line by default
	if chunklen == 0 {
		chunklen = 76
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

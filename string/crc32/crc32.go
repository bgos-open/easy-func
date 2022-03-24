package crc32

import "hash/crc32"

// crc32 — Calculates the crc32 polynomial of a string (https://www.php.net/manual/en/function.crc32.php)

func Do(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

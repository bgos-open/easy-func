package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// https://www.php.net/manual/en/function.md5
func Do(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

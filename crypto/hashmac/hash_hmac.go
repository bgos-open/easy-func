package hashmac

import (
	"crypto/hmac"
	"crypto/sha1" //nolint:gosec
	"encoding/hex"
)

// Do 使用 HMAC 方法生成带有密钥的哈希值 https://www.php.net/manual/zh/function.hash-hmac.php
func Do(input, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// https://www.php.net/manual/en/function.md5-file.php
func Do(filename string) string {
	if filename == "" {
		return fmt.Sprintf("文件路径不能为空")
	}
	path := fmt.Sprintf("./%s", filename)
	pFile, err := os.Open(path)
	if err != nil {
		return fmt.Sprintf("打开文件失败，filename=%v, err=%v", filename, err)
	}

	defer pFile.Close()
	md5Hash := md5.New()
	io.Copy(md5Hash, pFile)

	return hex.EncodeToString(md5Hash.Sum(nil))
}

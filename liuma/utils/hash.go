package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"liuma/exception/http_err"
)

// 哈希值计算
func CalculateHash(r io.Reader) (string, interface{}) {
	h := sha256.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return "", http_err.CalculateHashError()
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
	//return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", h.Sum(nil)))), nil
}


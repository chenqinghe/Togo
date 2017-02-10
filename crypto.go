package gophp

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	result := h.Sum(nil)
	return hex.EncodeToString(result), nil
}

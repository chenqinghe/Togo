package gophp

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	result := h.Sum(nil)
	return string(result), nil
}

package Togo

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"encoding/base64"
)

func Md5(s string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	result := h.Sum(nil)
	return hex.EncodeToString(result), nil
}

func Md5_file(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contentbyte, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	h:= md5.New()
	h.Write(contentbyte)
	return hex.EncodeToString(h.Sum(nil)), nil

}


func Base64_decode(str string) (string, error) {
	bt, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
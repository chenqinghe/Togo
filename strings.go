/*
字符串操作相关函数
*/
package Togo

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Explode(str, sep string) ([]string, error) {
	if sep == "" {
		return []string{str}, nil
	}
	return strings.Split(str, sep), nil
}

func HtmlSpecialCharsDecode(s string) string {
	s = strings.Replace(s, "&lt;", "<", -1)
	s = strings.Replace(s, "&gt;", ">", -1)
	s = strings.Replace(s, "&amp;", "&", -1)
	s = strings.Replace(s, `&quot;`, `"`, -1)
	s = strings.Replace(s, `&#039;`, `'`, -1)
	return s
}

func HtmlSpecialChars(s string) string {
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
	s = strings.Replace(s, "&", "&amp;", -1)
	s = strings.Replace(s, `"`, "&quot;", -1)
	s = strings.Replace(s, `'`, "&#039;", -1)
	return s

}

func Trim(s1, s2 string) string {
	return strings.Trim(s1, s2)
}

func Ltrim(s1, s2 string) string {
	return strings.TrimLeft(s1, s2)
}

func Rtrim(s1, s2 string) string {
	return strings.TrimRight(s1, s2)
}

func Strrev(s string) string {
	if s == "" {
		return s
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	length := len([]rune(s))
	var re []rune = make([]rune, length)
	i := 0
	for _, v := range s {
		re[length-i-1] = v
		i++
	}
	return string(re)
}

func Bin2hex(s string) (string, error) {
	i, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

func Hex2bin(s string) (string, error) {

	i, err := strconv.ParseInt(s, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

func Chr(ascii int) (string, error) {
	if ascii > 127 || ascii < 32 {
		return "", errors.New("invalid ascii code.")
	}
	var buf bytes.Buffer
	buf.Write([]byte{byte(ascii)})
	return buf.String(), nil
}

func BaseConvert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

func Base64Decode(str string) (string, error) {
	bt, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

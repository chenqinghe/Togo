/*
字符串操作相关函数
*/
package Togo

import (
	"bytes"
	"errors"
	"fmt"
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

func Chr(ascii int) (string, error) {
	if ascii > 127 || ascii < 32 {
		return "", errors.New("invalid ascii code.")
	}
	var buf bytes.Buffer
	buf.Write([]byte{byte(ascii)})
	return buf.String(), nil
}





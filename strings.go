package pfunc

import (
	"fmt"
	"strings"
)

func Explode(str, sep string) ([]string, error) {
	if sep == "" {
		return []string{str}, nil
	}
	return strings.Split(str, sep), nil
}

func Hex2bin(s string) (string, error) {
	return "", nil
}

func HtmlSpecialCharsDecode(s string) (string, error) {
	return "", nil
}

func HtmlSpecialChars(s string) string {
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
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

func Reverse(s string) string {
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

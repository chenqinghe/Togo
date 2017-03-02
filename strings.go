/*
字符串操作相关函数
*/
package Togo

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strings"
)

func Explode(str, sep string) ([]string, error) {
	if sep == "" {
		return []string{str}, nil
	}
	return strings.Split(str, sep), nil
}

func Htmlspecialchars_decode(s string) string {
	s = strings.Replace(s, "&lt;", "<", -1)
	s = strings.Replace(s, "&gt;", ">", -1)
	s = strings.Replace(s, "&amp;", "&", -1)
	s = strings.Replace(s, `&quot;`, `"`, -1)
	s = strings.Replace(s, `&#039;`, `'`, -1)
	return s
}

func Htmlspecialchars(s string) string {
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

/**
 *使用此函数将字符串分割成小块非常有用。
 *例如将 base64_encode() 的输出转换成符合 RFC 2045 语义的字符串。
 *它会在每 chunklen 个字符后边插入 end。(按byte计数，不是rune！)
 */
func Chunk_split(body string, chunklen int, end string) (str string, err error) {
	count := int(math.Ceil(float64(len(body) / chunklen)))
	var buf bytes.Buffer
	for i := 0; i < count-1; i++ { //防止访问越界
		start := i * chunklen
		_, err = buf.Write([]byte(body)[start : start+chunklen])
		if err != nil {
			break
		}
		_, err = buf.WriteString(end)
		if err != nil {
			break
		}
	}
	_, err = buf.Write([]byte(body)[chunklen*(count-1):])
	_, err = buf.WriteString(end)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

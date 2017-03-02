/*
 *此包内主要是php中数组相关的操作
 *本包内的数组均指map[string]interface{}
 */
package Togo

import "strings"

type Array map[string]interface{}

type Case int

const (
	CASE_UPPER Case = 1
	CASE_LOWER Case = 0
)

/**
计算数组的差集
在第一个array中但不在剩余array中的元素
*/
func Arraydiff(m ...Array) Array {
	return Array{}
}

func Array_change_key_case(arr Array, cs Case) Array {
	var tmp Array = Array{}
	for k, v := range arr {
		tmp[strings.ToUpper(k)] = v
	}
	return tmp
}

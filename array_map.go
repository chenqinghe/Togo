/*
 *此包内主要是php中关联数组相关的操作
 *本包内的数组均指map[string]interface{}
 */
package Togo

import "strings"

type ArrayMap map[string]interface{}

type Case int

const (
	CASE_UPPER Case = 1
	CASE_LOWER Case = 0
)

/**
计算数组的差集
在第一个ArrayMap中但不在剩余ArrayMap中的元素
*/
func Array_diff(m ...ArrayMap) ArrayMap {
	//todo
	return ArrayMap{}
}

func Array_change_key_case(arr ArrayMap, cs Case) ArrayMap {
	var tmp ArrayMap = ArrayMap{}
	if cs == CASE_LOWER {
		for k, v := range arr {
			tmp[strings.ToLower(k)] = v
		}
	} else if cs == CASE_UPPER {
		for k, v := range arr {
			tmp[strings.ToUpper(k)] = v
		}
	}
	return tmp
}

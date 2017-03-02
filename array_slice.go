/*
 *此包内主要是php中索引数组相关的操作
 *本包内的数组均指[]interface{}
 */
package Togo

import "fmt"

type ArraySlice []interface{}

func Array_chunk(input ArraySlice, size int) ArraySlice {
	length := len(input)
	count := int(Ceil(float64(length) / float64(size)))
	ret := make(ArraySlice, count)
	for i := 0; i < count-1; i++ {
		ret[i] = input[i*size : (i+1)*size]
		fmt.Println(ret)
	}
	ret[count-1] = input[size*(count-1):]
	return ret
}

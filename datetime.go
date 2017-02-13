/**
此包內是关于时间和日期的相关函数
 */

package togo

import "time"

func Date() {

}

func Time() int64 {
	return time.Now().Unix()
}

func Sleep(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}
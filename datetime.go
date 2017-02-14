/**
此包內是关于时间和日期的相关函数
*/
package Togo

import "time"

func Time() int64 {
	return time.Now().Unix()
}

func Sleep(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}

func CheckDate(month, day, year int) bool {
	//check month
	if month > 12 || month < 1 {
		return false
	}
	//check day
	if day < 1 {
		return false
	}
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 { //31 days
		if day > 31 {
			return false
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 { //30 days
		if day > 30 {
			return false
		}
	} else { // february
		if checkIfLeapYear(year) {
			if day > 29 {
				return false
			}
		}else {
			if day > 28 {
				return false
			}
		}
	}
	//check year
	if year < 1 || year > 32767 {
		return false
	}
	return true
}

func checkIfLeapYear(year int) bool {

	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func Date(str string ,timestamp int64)(string,error){

	return "",nil
}





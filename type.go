//这个包是关于数据类型转换的
package Togo

func Intval() {

}

func Strval() {

}

func Boolval(val interface{}) bool {

	switch v := val.(type) {
	case map[interface{}]interface{}:
		if len(v) != 0 {
			return true
		}
	case int, int8, int16, int32, int64:
		if v != 0 {
			return true
		}
	case bool:
		return v
	case complex64, complex128:
		if v != complex128(0) {
			return true
		}
	case float32, float64:
		if v != float64(0) {
			return true
		}
	case uint, uint8, uint16, uint32, uint64:
		if v != 0 {
			return true
		}
	default:
		return false
	}
	return false
}

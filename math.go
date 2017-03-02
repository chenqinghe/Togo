package Togo

import (
	"math"
	"math/cmplx"
	"strconv"
)

func Abs(n float64) float64 {
	return math.Abs(n)
}

func Acos(val complex128) complex128 {
	return cmplx.Acos(val)
}

func Acosh(val complex128) complex128 {
	return cmplx.Acosh(val)
}

func Asin(val complex128) complex128 {
	return cmplx.Asin(val)
}

func Asinh(val complex128) complex128 {
	return cmplx.Asinh(val)
}

func Atan(val complex128) complex128 {
	return cmplx.Atan(val)
}

func Atanh(val complex128) complex128 {
	return cmplx.Atanh(val)
}

func Base_convert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
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

func Bindec(bin string) (int64, error) {
	i, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		err = err.(*strconv.NumError).Err
	}
	return i, err
}

func Ceil(f float64)float64{
	return math.Ceil(f)
}

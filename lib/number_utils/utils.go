package number_utils

/*
   @Time: 2023/9/28 22:45
   @Author: david
   @File: number_utils
*/

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func ABS[T Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

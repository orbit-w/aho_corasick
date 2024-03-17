package math

import (
	"fmt"
	"testing"
)

/*
   @Author: orbit-w
   @File: math_test
   @2024 3月 周日 11:30
*/

func Test_PowerOf2(t *testing.T) {
	fmt.Println(PowerOf2(0))
	fmt.Println(PowerOf2(1))
	fmt.Println(PowerOf2(63))
	fmt.Println(PowerOf2(65))
	fmt.Println(PowerOf2(1023))
	fmt.Println(PowerOf2(1025))
}

func Test_GenericFls(t *testing.T) {
	fmt.Println(GenericFls(1023 - 1))
}

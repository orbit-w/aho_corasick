package misc

import (
	"fmt"
	"testing"
	"time"
)

/*
   @Author: orbit-w
   @File: misc_test
   @2024 3月 周日 11:27
*/

func Test_MSCast(t *testing.T) {
	start := time.Now().UnixNano()
	time.Sleep(time.Millisecond)
	fmt.Println(MSCast("test", start))
}

func Test_PrintMem(t *testing.T) {
	PrintMem()
}

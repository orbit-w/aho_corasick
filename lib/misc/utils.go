package misc

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
   @Author: orbit-w
   @File: utils
   @2023 10月 周一 18:56
*/

func UsedNano(start, count int64) string {
	used := time.Now().UnixNano() - start
	return strings.Join([]string{"used: ", strconv.FormatInt(used, 10), "ns , ", strconv.FormatInt(used/count, 10), " ns/op "}, " ")
}

func PrintMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	println(fmt.Sprintf("Sys = %v MiB, TotalAlloc = %v MiB, HeapAlloc = %v MiB, NumGC = %v, HeapObjs = %v, Goroutine = %v", bToMb(m.Sys),
		bToMb(m.TotalAlloc), bToMb(m.HeapAlloc), m.NumGC, m.HeapObjects, runtime.NumGoroutine()))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

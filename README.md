# AhoCorasickGo
AC automaton, the algorithm uses DAT(double array trie) to construct a finite-state machine and combined the concept of KMP to construct a mismatch pointer.
It is a kind of dictionary-matching algorithm that locates elements of a finite set of strings (the "dictionary") within an input text. It matches all strings simultaneously.

## Introduction:
Aho_corasick algorithm is better than traditional dict trie, for example:
* less memory usage, differ by one order of magnitude, only 10-20% of trie
* pattern matching is an order of magnitude faster

## Comparison:
Compared with traditional dict trie, Read a dictionary library with a vocabulary of 51W, 
contains Chinese character set, memory usage:
```
                        HeapAlloc(MiB)      HeapObjs        BuildTime(ms)
    
Trie                    454                 8506021         507

AhoCorasickGo           125                 215915          2100

```

## Usage:
```go
package main

import (
	"github.com/orbit-w/aho_corasick/aho_corasick_v2"
	"github.com/orbit-w/aho_corasick/lib/misc"
	"runtime"
	"time"
)

func main() {
	ac, _ := aho_corasick_v2.LoadDict("./../../data/filter.txt")
	in := []rune("sdwdhjsfq.cfsadwd")

	//text input replacement
	ac.Replace(in, '*')
	
	ac.FindAll(in)
	
}

```

## Benchmark:
```go

package benchmark

import (
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"testing"
)

var (
	text = "sdwdhomoeysadwdsdwdsdwD-¥¶¯sdd-0gd-0gswnch-uj? ch-uj?congs-anba-c-hoba-c-hosdwdaba-c-ho"
)

func Benchmark_ACFindAll(b *testing.B) {
	ac, _ := aho_corasick.LoadDict("path")
	in := []rune(text)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("FindAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ac.FindAll(in)
		}
	})
}

```
```shell
goos: darwin
goarch: arm64
pkg: github.com/orbit-w/aho_corasick/aho_corasick/benchmark
Benchmark_ACFindAll
Benchmark_ACFindAll/FindAll
Benchmark_ACFindAll/FindAll-8         	 1000000	      1101 ns/op
PASS

```
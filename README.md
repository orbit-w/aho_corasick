# AhoCorasickGo
AC automaton, the algorithm uses DAT(double array trie) to construct a finite-state machine and combined the concept of KMP to construct a mismatch pointer.
It is a kind of dictionary-matching algorithm that locates elements of a finite set of strings (the "dictionary") within an input text. It matches all strings simultaneously.

## Comparison:
硬件 MacAir Apple M2, 在文本处理能力和内存占用两方面跟传统字典 trie 做了比较:
```
场景1: 加载了本地的屏蔽词字典,不包含中文字符,词汇量为13W

# 内存占用
                        Trie                AhoCorasickGo
HeapAlloc/MiB           35                  19

HeapObjs                709750              172106

BuildTime/ms            41                  116

# 测试处理433长度的文本性能对比，硬件 MacAir Apple M2 (模式串："outlieroutliersoutliesoutlineoutlinedoutlinesoutliningoutliveoutliveddwoutliveroutliversoutlivesoutlivingoutlookoutlooksoutloveoutlovedoutlovesoutlovingoutlyingsdhwdhoutmansdhwdhoutmaneuverojhbdwoutmaneuveredshjdwdjoutmaneuveringsdjawhdoutmaneuversdwadadoutmannediwjdskjoutmanningkdfjjoutmanswundnoutmarchhjghcoutmarchedwsdoutmarcheswdwoutmarchinglksmcnskncwjfwajdmsdbwajdwakjdsjkdbaskdbakwdbkasbdakndbsnabdkwdbsandbsndbnv @@#dasdawd")
                        Trie                AhoCorasickGo
接口名称
FindAll                 53 μs/op            9 μs/op

******************************************************************************

在大幅提高文本处理性能前提下，内存使用更优; 当字典中词汇量更大, 单条词汇更长，包含中文字符情况下优势更加明显。

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
	ac, _ := aho_corasick_v2.LoadDict(dictDir)
	in := []rune("sdwdhjsfq.cfsadwd")

	//text input replacement
	ac.Replace(in, '*')
	
	ac.FindAll(in)
	
}

```
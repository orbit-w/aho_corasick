# AhoCorasickGo
AC automaton, the algorithm uses DAT(double array trie) to construct a finite-state machine and combined the concept of KMP to construct a mismatch pointer.
It is a kind of dictionary-matching algorithm that locates elements of a finite set of strings (the "dictionary") within an input text. It matches all strings simultaneously.

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/orbit-w/aho_corasick/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/orbit-w/aho_corasick/tree/master)
[![codecov](https://codecov.io/gh/orbit-w/aho_corasick/graph/badge.svg?token=6LYVYO4EAL)](https://codecov.io/gh/orbit-w/aho_corasick)
[![Go Report Card](https://goreportcard.com/badge/github.com/orbit-w/aho_corasick)](https://goreportcard.com/report/github.com/orbit-w/aho_corasick)

## Comparison:
硬件信息 MacAir Apple M2, 在文本处理能力和内存占用两方面跟传统字典 trie 做了比较:

* 场景1: 加载了本地的屏蔽词字典,不包含中文字符,词汇量为13W
    
    ```
    # 内存占用
                            Trie                AhoCorasickGo
    HeapAlloc/MiB           35                  14
    
    HeapObjs                709750              172106
    
    BuildTime/ms            41                  91
    
    # 接口测试1: 测试处理433长度的文本性能对比，结果如下：
    (模式串："outlieroutliersoutliesoutlineoutlinedoutlinesoutliningoutliveoutliveddwoutliveroutliversoutlivesoutlivingoutlookoutlooksoutloveoutlovedoutlovesoutlovingoutlyingsdhwdhoutmansdhwdhoutmaneuverojhbdwoutmaneuveredshjdwdjoutmaneuveringsdjawhdoutmaneuversdwadadoutmannediwjdskjoutmanningkdfjjoutmanswundnoutmarchhjghcoutmarchedwsdoutmarcheswdwoutmarchinglksmcnskncwjfwajdmsdbwajdwakjdsjkdbaskdbakwdbkasbdakndbsnabdkwdbsandbsndbnv @@#dasdawd")
                          Trie                AhoCorasickGo
    接口名称
    FindAll               53 μs/op            7 μs/op
  
    # 接口测试2: 测试处理36长度的文本性能对比，结果如下：
    (模式串："dswkjdajcsnccawdlsd;adksfucksdwdsdwd")
                          Trie                AhoCorasickGo
    接口名称
    Validate              0.85 μs/op          0.004 μs/op
  
    # 接口测试3: 关于文本替换，同样处理433长度的文本，同接口测试1的模式串相同，结果如下：
                          Trie                AhoCorasickGo
    接口名称
    Replace               10.4 μs/op          4.7 μs/op
    ```
* 场景2: 加载了本地的网络游戏屏蔽词字典,包含中文字符,片假名等,词汇量为82W

    ```
    # 内存占用
                            Trie                AhoCorasickGo
    HeapAlloc/MiB           454                 144
    
    HeapObjs                8505987             215912
    
    BuildTime/ms            497                 2200
    ```

## Summary：
* 随着模式串的长度 length 增长，double array高效搜索性能优势越明显；
* 在大幅提高文本处理性能前提下比较 double array和 trie 内存使用情况，当字典中词汇量小, 包含字符集少情况下, double array会产生冗余空间条目.
当字典中词汇量大，包含字符集多情况下，比如包含片假名、中文等，double array会保持更紧凑排列且保证极小内存使用开销，优势更加明显.
引用原论文的话；
```
‘The number of redundant entries of the double-array grows for small sets of keys,
but the number for large sets can keep an extremely small value. In order to build
a more compact dictionary for small sets of keys, the remapping of characters on
the basis of their frequency, statistically, becomes necessary. In this implementation,
other kinds of characters (Katakana, Chinese, etc.) can be used. Nevertheless, it is
better to treat a multi-byte character as one-byte by one-byte due to an offset based
on a large numerical value makes the size of the double-array grow, that is to say,
the double-array has many available, or redundant, entries.’

感谢 JUN-ICHI AOE 论文
```

## Usage:
```go
package main

import (
	"github.com/orbit-w/aho_corasick/aho_corasick"
)

func main() {
    dictDir := "./data/en/dict.txt"
	
    var ac aho_corasick.AC

    if err := ac.LoadDict("./data/en/dict.txt"); err != nil {
      panic(err)
    }

    pattern := "sdwdhjsfq.cfsadwd"
  
    //text input replacement
    ac.Replace(pattern, '*')
  
    ac.FindAll([]rune(pattern))
	
}

```
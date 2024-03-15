package test

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"github.com/orbit-w/aho_corasick/lib/misc"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

/*
   @Author: orbit-w
   @File: mem_test
   @2023 10月 周日 15:15
*/

func Test_ACLoadEnDict(t *testing.T) {
	start := time.Now().UnixNano()
	var ac aho_corasick.AC

	err := ac.LoadDict(enDictDir)
	//ac, err := aho_corasick.LoadDict("./../../data/SenstiveWord.txt")
	assert.NoError(t, err)
	fmt.Println(ac.Cap() * 8 * 3 / 1024 / 1024)
	pattern := "sdwdhjsfq.cfsadwd"
	re := ac.Replace(pattern, '*')
	fmt.Println(misc.MSCast("AC", start))
	fmt.Println(re)
	runtime.GC()
	misc.PrintMem()
	ac.FindAll([]rune("sdwdhomoeysadwd"))
	runtime.GC()
	misc.PrintMem()
	time.Sleep(time.Second * 10)
	ac.Replace(pattern, '*')
	misc.PrintMem()
}

func Test_ACLoadCnDict(t *testing.T) {
	start := time.Now().UnixNano()
	var ac aho_corasick.AC
	err := ac.LoadDict(cnDictDir)
	assert.NoError(t, err)
	fmt.Println(ac.Cap())
	pattern := "太阳翼漫无边际测试用例"
	re := ac.Replace(pattern, '*')
	fmt.Println(misc.MSCast("AC", start))
	fmt.Println(re)
	runtime.GC()
	patterns := ac.FindAll([]rune("太阳翼漫无边际测试用例"))
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
	misc.PrintMem()
}

func Test_TrieLoadEnDict(t *testing.T) {
	start := time.Now().UnixNano()
	filter := sensitive.New()
	err := filter.LoadWordDict(enDictDir)
	assert.NoError(t, err)
	fmt.Println(filter.Replace("sdwdhjsfq.cfsadwd", '*'))
	fmt.Println(misc.MSCast("Trie", start))
	runtime.GC()
	filter.FindAll("sdwdhomoeysadwd")
	misc.PrintMem()
}

func Test_TrieLoadCnDict(t *testing.T) {
	start := time.Now().UnixNano()
	filter := sensitive.New()
	err := filter.LoadWordDict(cnDictDir)
	assert.NoError(t, err)
	fmt.Println(filter.Replace("太阳翼漫无边际测试用例", '*'))
	fmt.Println(misc.MSCast("Trie", start))
	runtime.GC()
	filter.FindAll("太阳翼漫无边际测试用例")
	misc.PrintMem()
}

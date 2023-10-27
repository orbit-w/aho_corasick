package test

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"github.com/orbit-w/aho_corasick/aho_corasick_v2"
	"github.com/orbit-w/aho_corasick/lib/misc"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

/*
   @Time: 2023/8/22 08:20
   @Author: david
   @File: dat_test
*/

var (
	text    = "outlearningsdwdsdoutgnawedsdwdsdad"
	dictDir = "./../../data/en/dict.txt"
)

// h: 104, e: 101, s: 115, r: 114, i: 105
func TestAC_Fail(t *testing.T) {
	ks := aho_corasick_v2.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := aho_corasick_v2.New(ks)
	ac.Print()
	fmt.Println(rune('C'))
}

func TestAC_MultiPatternSearch(t *testing.T) {
	ks := aho_corasick_v2.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	ac := aho_corasick_v2.New(ks)
	input := []rune("ahishers")
	patterns := ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}

	fmt.Println("===================================================")
	//中文模式串匹配
	input = []rune("听说独立日日这部电影是美国人拍的")
	ks = aho_corasick_v2.StrKeySlice{
		[]rune("独立"),
		[]rune("独立日"),
	}
	ac = aho_corasick_v2.New(ks)
	patterns = ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
}

func Test_ACLoad(t *testing.T) {
	start := time.Now().UnixNano()
	//ac, err := aho_corasick_v2.LoadDict("./../../data/SenstiveWord.txt")
	ac, err := aho_corasick_v2.LoadDict(dictDir)
	assert.NoError(t, err)
	fmt.Println(ac.Cap())
	in := []rune("sdwdhjsfq.cfsadwd")
	ac.Replace(in, '*')
	fmt.Println(misc.MSCast("AC", start))
	fmt.Println(string(in))
	runtime.GC()
	ac.FindAll([]rune("sdwdhomoeysadwd"))
	misc.PrintMem()
}

func Test_ACFindAll(t *testing.T) {
	ac, err := aho_corasick_v2.LoadDict(dictDir)
	assert.NoError(t, err)
	in := []rune(text)
	res := ac.FindAll(in)
	for i := range res {
		r := res[i]
		fmt.Println(string(r.Pattern))
	}
}

func Test_ACReplace(t *testing.T) {
	ks := aho_corasick_v2.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}
	ac := aho_corasick_v2.New(ks)
	input := []rune("ahisherssadwdshershis")
	for _, v := range ac.FindAll(input) {
		fmt.Println(string(v.Pattern))
	}
	ac.Replace(input, '*')
	fmt.Println(string(input))
	assert.Equal(t, "a*******sadwd********", string(input))
}

func Test_Replace(t *testing.T) {
	filter := sensitive.New()
	err := filter.LoadWordDict(dictDir)
	assert.NoError(t, err)
	str1 := filter.Replace(text, '*')
	fmt.Println(filter.Replace(text, '*'))

	ac, err := aho_corasick_v2.LoadDict(dictDir)
	assert.NoError(t, err)
	in := []rune(text)
	ac.Replace(in, '*')
	fmt.Println(string(in))
	assert.Equal(t, str1, string(in))
}

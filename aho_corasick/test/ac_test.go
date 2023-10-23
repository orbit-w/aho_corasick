package test

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"github.com/orbit-w/aho_corasick/lib/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
   @Time: 2023/8/22 08:20
   @Author: david
   @File: dat_test
*/

// h: 104, e: 101, s: 115, r: 114, i: 105
func TestAC_Fail(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := aho_corasick.New(ks)
	ac.Print()
	fmt.Println(rune('C'))
}

func TestAC_MultiPatternSearch(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	ac := aho_corasick.New(ks)
	input := []rune("ahishers")
	patterns := ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}

	fmt.Println("===================================================")
	//中文模式串匹配
	input = []rune("听说独立日日这部电影是美国人拍的")
	ks = aho_corasick.StrKeySlice{
		[]rune("独立"),
		[]rune("独立日"),
	}
	ac = aho_corasick.New(ks)
	patterns = ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
}

func Test_ACLoadAndSearch(t *testing.T) {
	ac, err := aho_corasick.LoadDict("./../../data/filter_dict.txt")
	assert.NoError(t, err)
	fmt.Println(ac.Cap())
	misc.PrintMem()
}

func Test_ACReplace(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}
	ac := aho_corasick.New(ks)
	input := []rune("ahisherssadwdshershis")
	for _, v := range ac.FindAll(input) {
		fmt.Println(string(v.Pattern))
	}
	ac.Replace(input, '*')
	fmt.Println(string(input))
	assert.Equal(t, "a*******sadwd********", string(input))
}

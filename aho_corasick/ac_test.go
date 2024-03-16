package aho_corasick

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
   @Time: 2023/8/22 08:20
   @Author: orbit-w
   @File: dat_test
*/

var (
	text      = "outlearningsdwdsdoutgnawedsdwdsdad"
	enDictDir = "./../data/en/dict.txt"
	cnDictDir = "./../data/cn/dict.txt"
)

// h: 104, e: 101, s: 115, r: 114, i: 105
func TestAC_Fail(t *testing.T) {
	ks := StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := New(ks)
	ac.Print()
}

func TestAC_MultiPatternSearch(t *testing.T) {
	ks := StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	ac := New(ks)
	input := []rune("ahishers")
	patterns := ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}

	fmt.Println("===================================================")
	//中文模式串匹配
	input = []rune("听说独立日日这部电影是美国人拍的")
	ks = StrKeySlice{
		[]rune("独立"),
		[]rune("独立日"),
	}
	ac = New(ks)
	patterns = ac.FindAll(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
}

func Test_ACFindAll(t *testing.T) {
	var ac AC
	err := ac.LoadDict(enDictDir)
	assert.NoError(t, err)
	in := []rune(text)
	res := ac.FindAll(in)
	for i := range res {
		r := res[i]
		fmt.Println(string(r.Pattern))
	}
}

func Test_ACReplace(t *testing.T) {
	ks := StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}
	ac := New(ks)
	pattern := "ahisherssadwdshershis"
	input := []rune(pattern)
	for _, v := range ac.FindAll(input) {
		fmt.Println(string(v.Pattern))
	}
	str := ac.Replace(pattern, '*')
	fmt.Println(str)
	assert.Equal(t, "a*******sadwd***rs***", str)
}

func Test_Replace(t *testing.T) {
	filter := sensitive.New()
	err := filter.LoadWordDict(enDictDir)
	assert.NoError(t, err)
	fmt.Println(filter.FindAll(text))
	str1 := filter.Replace(text, '*')
	fmt.Println(str1)
	filter.FindAll(text)

	var ac AC

	err = ac.LoadDict(enDictDir)
	assert.NoError(t, err)
	//"outlearningsdwdsdoutgnawedsdwdsdad"
	str2 := ac.Replace(text, '*')
	fmt.Println(str2)
	assert.Equal(t, str1, str2)

	//测试ReplaceAll接口，期望结果：‘************dwds***********dwds***’
	str2 = ac.ReplaceAll(text, '*')
	fmt.Println(str2)
}

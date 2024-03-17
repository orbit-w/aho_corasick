package aho_corasick

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*
   @Time: 2023/10/17 16:14
   @Author: orbit-w
   @File: dat_test
*/

func Test_DATLoadFile(t *testing.T) {
	var dat DAT
	err := dat.LoadDict(enDictDir)
	assert.NoError(t, err)

	assert.True(t, dat.Find([]rune("whity")))
}

func TestDAT_Find(t *testing.T) {
	ks := StrKeySlice{
		[]rune("hf"),
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}
	sort.Sort(ks)
	trie := new(Trie)
	trie.Build(ks)
	dat := new(DAT)
	dat.Build(trie)

	fmt.Println(dat.Find([]rune("herss")))
	fmt.Println(dat.Find([]rune("hers")))
	fmt.Println(dat.Find([]rune("his")))
}

func TestAC_AutomationInit(t *testing.T) {
	keywords := [][]string{
		{"清", "华"},
		{"清", "华", "大", "学"},
		{"清", "新"},
		{"中", "华"},
		{"华", "人"},
	}
	ks := StrKeySlice{}

	for _, keyword := range keywords {
		dk := make(StrKey, 0)
		for _, k := range keyword {
			dk = append(dk, toRune(k))
		}
		ks = append(ks, dk)
	}
	sort.Sort(ks)
	trie := new(Trie)
	trie.Build(ks)
	dat := new(DAT)
	dat.Build(trie)
	dat.Print()
}

func Test_Print(t *testing.T) {
	fmt.Println('A')
	fmt.Println('C')
	fmt.Println('Z')
	fmt.Println('D')
	fmt.Println('F')

	ks := StrKeySlice{
		[]rune("AC"),
		[]rune("ACE"),
		[]rune("ACFF"),
		[]rune("AD"),
		[]rune("CD"),
		[]rune("CF"),
		[]rune("ZQ"),
	}
	sort.Sort(ks)
	for i := range ks {
		s := string(ks[i])
		fmt.Println(s)
	}
}

func Test_DATLenAndCap(t *testing.T) {
	keywords := [][]string{
		{"清", "华"},
		{"清", "华", "大", "学"},
		{"清", "新"},
		{"中", "华"},
		{"华", "人"},
	}
	ks := StrKeySlice{}

	for _, keyword := range keywords {
		dk := make(StrKey, 0)
		for _, k := range keyword {
			dk = append(dk, toRune(k))
		}
		ks = append(ks, dk)
	}
	sort.Sort(ks)
	trie := new(Trie)
	trie.Build(ks)
	dat := new(DAT)
	dat.Build(trie)
	fmt.Println(dat.Length())
	fmt.Println(dat.Cap())
}

func toRune(s string) rune {
	switch s {
	case "清":
		return 1
	case "华":
		return 2
	case "大":
		return 3
	case "学":
		return 4
	case "新":
		return 5
	case "中":
		return 6
	case "人":
		return 7
	}
	return 0
}

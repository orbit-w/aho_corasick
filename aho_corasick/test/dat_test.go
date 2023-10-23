package test

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"sort"
	"testing"
)

/*
   @Time: 2023/10/17 16:14
   @Author: david
   @File: dat_test
*/

func TestDAT_Find(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}
	sort.Sort(ks)
	trie := new(aho_corasick.Trie)
	trie.Build(ks)
	dat := new(aho_corasick.DAT)
	dat.Build(trie)
	dat.Print()

	fmt.Println(dat.Find([]rune("herss")))
	fmt.Println(dat.Find([]rune("hers")))
}

func TestAC_AutomationInit(t *testing.T) {
	keywords := [][]string{
		{"清", "华"},
		{"清", "华", "大", "学"},
		{"清", "新"},
		{"中", "华"},
		{"华", "人"},
	}
	ks := aho_corasick.StrKeySlice{}

	for _, keyword := range keywords {
		dk := make(aho_corasick.StrKey, 0)
		for _, k := range keyword {
			dk = append(dk, toRune(k))
		}
		ks = append(ks, dk)
	}
	sort.Sort(ks)
	trie := new(aho_corasick.Trie)
	trie.Build(ks)
	dat := new(aho_corasick.DAT)
	dat.Build(trie)
	dat.Print()
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

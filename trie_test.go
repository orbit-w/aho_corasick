package aho_corasick

import (
	"fmt"
	"sort"
	"testing"
)

/*
   @Author: orbit-w
   @File: trie_test
   @2024 3月 周日 13:59
*/

func Test_TrieDFS(t *testing.T) {
	ks := StrKeySlice{
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

	ac := new(AC)

	length := dat.Length()
	ac.cap = length
	ac.base = make([]int, length)
	copy(ac.base, dat.base)
	ac.check = make([]int, length)
	copy(ac.check, dat.check)

	ac.output = make([]Output, length)

	ac.fail = make([]int, ac.cap)
	trie.DFS(func(node *Node) (stop bool) {
		fmt.Println(node.code)
		return
	})
}

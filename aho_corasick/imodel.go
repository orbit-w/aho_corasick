package aho_corasick

/*
   @Author: orbit-w
   @File: model
   @2023 10月 周五 18:55
*/

type IAhoCorasick interface {
	Cap() int
	Replace(input []rune, repl rune)
	FindAll(input []rune) []Result
	Validate(input []rune) bool
	Print()
}

type IDat interface {
	Build(trie *Trie)
	Find(keyword []rune) bool
	Length() int
	Cap() int
	State(s int) int
	Empty(s int) bool
}

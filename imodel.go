package aho_corasick

/*
   @Author: orbit-w
   @File: model
   @2023 10月 周五 18:55
*/

type IAhoCorasick interface {
	Cap() int
	Replace(pattern string, repl rune) string
	ReplaceAll(pattern string, repl rune) string
	FindAll(input []rune) []Result
	Validate(input []rune) bool
	Print()
}

type IDat interface {
	Build(trie *Trie)
	Find(keyword []rune) bool
	Length() int
	Cap() int
	Empty(s int) bool
}

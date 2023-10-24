package benchmark

import (
	"github.com/importcjj/sensitive"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
   @Author: orbit-w
   @File: benchmark_FinaAll
   @2023 10月 周二 11:49
*/

var (
	text = "sdwdhomoeysadwdsdwdsdwD-¥¶¯sdd-0gd-0gswnch-uj? ch-uj?congs-anba-c-hoba-c-hosdwdaba-c-ho"
)

func Benchmark_ACFindAll(b *testing.B) {
	ac, err := aho_corasick.LoadDict("./../../data/filter_dict.txt")
	assert.NoError(b, err)
	in := []rune(text)

	b.ReportAllocs()
	b.ResetTimer()
	b.Run("FindAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ac.FindAll(in)
		}
	})
}

func Benchmark_ACReplace(b *testing.B) {
	ac, err := aho_corasick.LoadDict("./../../data/filter_dict.txt")
	assert.NoError(b, err)
	in := []rune(text)

	b.ReportAllocs()
	b.ResetTimer()
	b.Run("Replace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ac.Replace(in, '*')
		}
	})
}

func Benchmark_TrieFindAll(b *testing.B) {
	filter := sensitive.New()
	err := filter.LoadWordDict("./../../data/filter_dict.txt")
	assert.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("FindAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter.FindAll(text)
		}
	})
}

func Benchmark_TrieReplace(b *testing.B) {
	filter := sensitive.New()
	err := filter.LoadWordDict("./../../data/filter_dict.txt")
	assert.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("Replace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = filter.Replace(text, '*')
		}
	})
}
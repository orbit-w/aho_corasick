package benchmark

import (
	"github.com/importcjj/sensitive"
	"github.com/orbit-w/aho_corasick/aho_corasick_v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
   @Author: orbit-w
   @File: benchmark_FinaAll
   @2023 10月 周二 11:49
*/

var (
	text = "sdwdhomoeys秀发发布周sdwD-¥¶¯sdd-0gd-0gswnch-uj? ch-uj?陶瓷展-辉煌夺目-c-牵强附会-c-hosdwdaba-c-ho"
)

func Benchmark_ACFindAll(b *testing.B) {
	ac, err := aho_corasick_v2.LoadDict("./../../data/filter.txt")
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
	ac, err := aho_corasick_v2.LoadDict("./../../data/filter.txt")
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
	err := filter.LoadWordDict("./../../data/filter.txt")
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
	err := filter.LoadWordDict("./../../data/filter.txt")
	assert.NoError(b, err)
	b.ReportAllocs()
	b.ResetTimer()
	b.Run("Replace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = filter.Replace(text, '*')
		}
	})
}

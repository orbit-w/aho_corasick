package test

import (
	"fmt"
	"github.com/importcjj/sensitive"
	"github.com/orbit-w/aho_corasick/lib/misc"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

/*
   @Author: orbit-w
   @File: trie_test
   @2023 10月 周一 22:00
*/

func Test_Trie(t *testing.T) {
	start := time.Now().UnixNano()
	filter := sensitive.New()
	err := filter.LoadWordDict("./../../data/filter_dict.txt")
	assert.NoError(t, err)
	runtime.GC()
	fmt.Println(filter.Replace("sdwdhomoeysadwd", '*'))
	fmt.Println(misc.MSCast("Trie", start))
	misc.PrintMem()
}

func Test_TrieFindAll(t *testing.T) {
	filter := sensitive.New()
	err := filter.LoadWordDict("./../../data/filter_dict.txt")
	assert.NoError(t, err)
	in := "sdwdhomoeysadwdsdwdsdwD-¥¶¯sdd-0gd-0gswnch-uj? ch-uj?"
	fmt.Println(filter.FindAll(in))
}

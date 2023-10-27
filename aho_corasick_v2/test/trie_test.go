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
	err := filter.LoadWordDict(dictDir)
	assert.NoError(t, err)
	fmt.Println(filter.Replace("sdwdhjsfq.cfsadwd", '*'))
	fmt.Println(misc.MSCast("Trie", start))
	runtime.GC()
	filter.FindAll("sdwdhomoeysadwd")
	misc.PrintMem()
}

func Test_TrieFindAll(t *testing.T) {
	filter := sensitive.New()
	err := filter.LoadWordDict(dictDir)
	assert.NoError(t, err)
	fmt.Println(filter.FindAll(text))
}

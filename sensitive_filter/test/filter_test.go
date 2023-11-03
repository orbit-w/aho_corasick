package test

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/sensitive_filter"
	"testing"
	"time"
)

/*
   @Author: orbit-w
   @File: filter_test
   @2023 11月 周五 17:06
*/

func Test_Reload(t *testing.T) {
	loader := sensitive_filter.NewLoader()
	_ = loader.LoadDict("./../../data/en/dict.txt")

	f := sensitive_filter.NewFilter(loader)

	pattern := "jjnzb"
	//jjnzb
	fmt.Println(f.Validate(pattern))

	loader = sensitive_filter.NewLoader()
	_ = loader.LoadDict("./../../data/en/dict.txt")
	loader.Merge([]string{"jjnzb"})
	f.ReBuild(loader)

	time.Sleep(time.Second * 10)
	fmt.Println(f.Validate(pattern))
}

package sensitive_filter

import (
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"regexp"
	"sync/atomic"
	"unsafe"
)

/*
   @Author: orbit-w
   @File: filter
   @2023 11月 周五 11:12
*/

type ISensitiveFilter interface {
	ReBuild(loader *Loader)
	Validate(pattern string) bool
	Replace(pattern string, repl rune) string
}

type SensitiveFilter struct {
	*aho_corasick.AC
	noise *regexp.Regexp
}

func NewFilter(loader *Loader) ISensitiveFilter {
	ac := new(aho_corasick.AC)
	loader.build(ac)
	return &SensitiveFilter{
		AC:    ac,
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
}

func (ins *SensitiveFilter) ReBuild(loader *Loader) {
	go func() {
		ac := new(aho_corasick.AC)
		loader.build(ac)
		p := (*unsafe.Pointer)(unsafe.Pointer(&ins.AC))
		atomic.StorePointer(p, unsafe.Pointer(ac))
	}()
}

func (ins *SensitiveFilter) Validate(pattern string) bool {
	cp := ins.RemoveNoise(pattern)
	return ins.AC.Validate([]rune(cp))
}

func (ins *SensitiveFilter) RemoveNoise(pattern string) string {
	return ins.noise.ReplaceAllString(pattern, "")
}

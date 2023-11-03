# SensitiveFilter

敏感词过滤器, AhoCorasickGo 应用的场景实例, 不支持动态增加/删除；如果需要调整词库中的数据，则需要异步重新构建词库。
```go
type ISensitiveFilter interface {
	ReBuild(loader *Loader)
	Validate(pattern string) bool
	Replace(pattern string, repl rune) string
}
```

## Usage:
```go
package main

import (
	"github.com/orbit-w/aho_corasick/sensitive_filter"
)

func main() {
	loader := sensitive_filter.NewLoader()
	_ = loader.LoadDict("./data/en/dict.txt")

	f := sensitive_filter.NewFilter(loader)

	pattern := "sdwdhjsfq.cfsadwd"
	f.Replace(pattern, '*')

	f.Validate(pattern)
}

```
# Golang 数据结构扩展：集合(set)

## 引入 & 使用

```bash
go get -u github.com/uberate/gset
```

## 样例

```go
package test

import "github.com/Uberate/gset"

func test() {
	set := gset.FromArray([]String{"u", "b", "e", "r", "a", "t", "e"})
	set.Delete("u")
}
```
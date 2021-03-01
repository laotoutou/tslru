# tslru

一个线程安全的、支持高并发的lru库

## 安装

```
go get github.com/laotoutou/tslru
```

## 使用

```
// 最大长度为4, 分片个数=4
c, err := lru2.NewTSCache(5, 4)
if err != nil {
	fmt.Println(err)
	return
}

c.Add("1", 1)

_, _ = c.Get("1")
```

## 博客

[LRU并发Get性能优化](https://laotoutou.github.io/posts/golang-lru-bing-fa-get-xingneng-youhua.html)

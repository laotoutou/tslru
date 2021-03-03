# tslru

一个线程安全的、支持高并发的lru库

## 安装

```
go get github.com/laotoutou/tslru
```

## 使用

```
lru, err := tslru.NewLRU(3)
if err != nil {
	fmt.Errorf("error=%s", err)
	return
}

lru.Add(1, 1)
lru.Add(2, 2)
lru.Add(3, 3)
lru.Add(4, 4)

fmt.Println(lru.GetOldest())
fmt.Println(lru.Keys())
fmt.Println(lru.Get(2))
fmt.Println(lru.Keys())
```

## 博客

[LRU并发Get性能优化](https://laotoutou.github.io/posts/golang-lru-bing-fa-get-xingneng-youhua.html)

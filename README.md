# tslru
一个线程安全的、支持高并发的lru库

## 安装

```
go get github.com/laotoutou/tslru
```

## 使用

```
cache := lru.NewTSCache(1024, 4)
```

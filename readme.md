## 热点缓存
github.com/hashicorp/golang-lru(1.0.2)、groupcache

## 应用进程资源消耗
net/http/pprof 排查应用进程资源消耗
https://zhuanlan.zhihu.com/p/396363069
curl "http://127.0.0.1:6060/debug/pprof/allocs?seconds=180" -o hello_process_alloc_120.conf 可获取120s内的栈内存消耗情况
curl "http://127.0.0.1:6060/debug/pprof/heap?seconds=300" -o hello_process_heap_300.conf 可获取300s内的堆内存消耗情况
yum install graphviz
go tool pprof -inuse_space -cum -svg hello_process_alloc_120.conf > alloc_300.svg

defer debug.FreeOSMemory() 能手动触发内存回收，但对CPU消耗较大




## 调试
go get github.com/hashicorp/golang-lru@v1.0.2
报错 cannot find module providing package github.com/hashicorp/golang-lru: working directory is not part of a module
go mod init golang-lru
go mod edit -require github.com/hashicorp/golang-lru@v1.0.2
# LRU
最近最少使用(最长时间)淘汰算法（Least Recently Used）。LRU是淘汰最长时间没有被使用的页面

# LFU
最不经常使用(最少次)淘汰算法（Least Frequently Used）。LFU是淘汰一段时间内，使用次数最少的页面

# golang-lru
![img.png](images/img.png)
![img_1.png](images/img_1.png)
![img.png](images/img_2.png)

golang-lru是一个基于双向链表实现的线程安全LRU缓存工具，若数据已经在缓存中，将其移到队首，并返回结果。若数据不在缓存中，将新记录添加到队首。若缓存超长，清理队尾缓存数据。


# evict
test: lru_evict.go
# not evict
test: lru_not_evict.go
# thread safe
test: lru_thread_safe.go
# accounts memory usage
lru_memory_usage.go

## accounts占用内存情况
# 每个用户40个资产
accounts数量->占用内存

1000000->     7666.950 M

100000->       798.119 M

10000->       82.471 M

1000->         9.875 M

# 每个用户20个资产
accounts数量->占用内存

1000000->     4030.459 M

100000->       446.180 M

10000->       46.937 M

1000->         9.625 M


## nfts占用内存情况
nfts数量->占用内存

10000000->    3387.616 M

1000000->      362.268 M

100000->        38.187 M

10000->          9.312 M

1000->           0 M

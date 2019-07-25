# redis通信协议

自定义文本协议

# 数据类型

string list hash set  sorted set 

hayperlog geo bit

# 数据淘汰策略

maxmemory

noeviction： 拒绝写入请求
volatile-lru: 淘汰设置了过期时间，最少使用的被淘汰
volatile-ttl: key的剩余ttl越短越先淘汰
近似LRU链表算法

# 一个字符串能存储多大值
512M

# 集群化方案
1. codis hash代理

2. cluster 不是一致性hash而是hash槽

# 集群方案什么情况会导致整个集群不可用

A，B，C三个节点没有复制模型的情况下如果节点B失败了那么整个集群就会缺少部分槽失败

# 设置密码与验证密码

# redis hash槽的概念

redis自带集群没有使用一致性hash, 而是引入了hash槽的概念，redis集群一共有16384个hash槽, 每个key通过CRC校验后对16384取模
来决定放置那个槽，集群的每个节点负责一部分槽点

# redis 集群的主从复制模型

为了使部分节点失败或者大部分节点无法通信的情况下集群仍然可用，所以集群使用了主从复制模型，每个节点都会有N-1个复制品

# redis 集群会有写操作丢失吗？ 为什么

redis并不能保证数据的强一致性，这意味着集群在特定条件下可能会丢失

# redis 集群间是怎样复制的 

异步复制

# redis 集群最大节点个数

16384

# redis 集群如何做数据库选择

目前集群不支持

# redis 中的管道有什么用


# 怎样理解redis 事务

# redis 如何做内存优化

尽量使用散列表

# redis 内存回收进程如何工作

# 内存回收算法

# redis 如何把做大量数据插入

pipe mode

# 为什么要做redis 分区

分区可以让redis管理更大的内存，redis将可以使用所有机器的内存。如果没有分区你最多使用一台机器的内存。分区使redis的计算能力通过简单
的增加计算机得到成倍的提升，redis的网络带宽也会随计算机的网卡的增加成倍增长

# redis 分区实现方案， 有哪些缺点？

# redis 持久化数据和缓存扩容怎么做？

# 分布式 redis 是前期做还是后期规模上来了再做

# 有哪些方式降低Redis 内存

# Redis 常见性能问题与解决方案

# 修改配置不重启 redis会生效吗









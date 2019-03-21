# mysql日志系统

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/mysql/mysql%E6%97%A5%E5%BF%97%E7%B3%BB%E7%BB%9F.jpeg'>
</p>

# 分类

- redo log 重做日志  InnoDB引擎独有的日志
- binlog 归档日志

# WAL 
先写日志再写磁盘

# redo log

- InnoDB引擎特有
- 物理日志，记录的是在某个数据页做了什么
- 循环写的，空间固定会用完继续从头开始写覆盖之前的

InnoDB的redo log是固定大小的，可以配置一组固定数量的文件，文件大小固定。从头开始写，写到末尾从开始重写<br><br>
redo log使数据库发生异常重启的时候之前提交的记录都不会丢失，这个能力我们叫 crash-safe

# binlog

- mysql服务层自带
- 逻辑日志，记录的是这个语句的原始逻辑
- 追加写入，不会覆盖之前的

```
innodb_flush_log_at_trx_commit = 1 // 这个参数设置为1的时候表示每次事物的redo log都可以持久化到磁盘。
sync_binlog = 1    // 每次事物的binlog都持久化到磁盘
```

# 两阶段提交

redo log 与 binlog 都可以用于表示事物的提交状态，两阶段提交让这两个状态逻辑上保持一致。









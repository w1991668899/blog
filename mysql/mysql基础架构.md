# mysql基础架构

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/mysql/mysql%E5%9F%BA%E6%9C%AC%E6%9E%B6%E6%9E%84.jpg'>
</p>

**参考:**

-   [mysql文档](https://dev.mysql.com/doc/refman/8.0/en/)
- 《高性能mysql》
- 《Mysql技术内幕》

# server层

- 连接器 管理连接，权限验证
- 缓存器 mysql8.0后面版本已经删除,不建议开启
- 分析器 词法分析，语法分析
- 优化器 执行计划生成，索引选择
- 执行器 操作引擎，返回结果

## 连接器

```
wt-001% mysql -h 127.0.0.1 -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 805
Server version: 8.0.13 MySQL Community Server - GPL

Copyright (c) 2000, 2018, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 

```

用户名验证通过后，连接器回去权限表查出权限，之后这个连接里面的操作都依赖此时读到的权限。这就意味着一旦连接成功后即使管理员对该用户更改权限也不影响已经连接的用户。只有重新连接才会使用新的权限。<br><br>

可以使用 `show pricesslist` 查看该连接状态，如下：

```
mysql> show processlist;
+-----+-----------------+------------------+----------------+---------+---------+------------------------+------------------+
| Id  | User            | Host             | db             | Command | Time    | State                  | Info             |
+-----+-----------------+------------------+----------------+---------+---------+------------------------+------------------+
|   4 | event_scheduler | localhost        | NULL           | Daemon  | 5611759 | Waiting on empty queue | NULL             |
| 804 | root            | 172.17.0.1:36476 | ttsing_service | Sleep   |    3404 |                        | NULL             |
| 805 | root            | 172.17.0.1:36490 | NULL           | Query   |       0 | starting               | show processlist |
+-----+-----------------+------------------+----------------+---------+---------+------------------------+------------------+
3 rows in set (0.05 sec)

mysql> 

```

**短链接**<br><br>
每次执行完sql后就断开连接，因为基于tcp建立连接开销比较大，相对也复杂所以建议大家使用长连接

**长连接**<br><br>
使用长连接后当连接量大的时候内存增长会很快。因为mysql在执行过程中临时使用的内存是管理在连接对象里面的。这些资源在断开的时候才释放.<br><br>
长连接使用内存过大会被系统强行杀掉。

我们在使用mysql连接池可以使用一下方式避免占用内存过大的问题：

- 定期断开长连接，之后使用再重新连接。可以通过 `wait_timeout` 控制默认连接时间, 一般连接迟启动配置参数中都有此参数，放到配置文件中来控制即可。
- mysql5.7及更高版本可以通过 `mysql_reset_connection` 重新初始化连接。这个过程不需要重连和权限验证，但会将连接恢复到刚刚创建完成时的状态。

## 缓存器

之前执行过的语句会以 k-v 形式被缓存在内存中。k 就是查询的语句，v 是查询结果。

**注意**
个人建议不要开启缓存，因为开启缓存弊大于利。缓存的失效是很频繁的，只要对一个表更新，该表相关的所有缓存都会被清空。<br><br>
将参数 `query_cache_type=DEMAND` 默认将会对所有sql不查询缓存.
这里的缓存是可以在sql中按需使用的,如下:
```
select SQL_CACHE * from t where id = 100;  //该语句在执行的时候会先查询缓存
```

mysql8.0版本中已经将缓存去掉。

## 分析器

分析器先做词法分析，mysql要分析我们输入的sql字符串中的字符分别代表什么，然后做语法分析看是否满足 SQL语句规则。一般语法错误就是从这里返回的。<br><br>
内建分析树，对其语法检查，先 from, 在 on, 再 join, 再 where, 检查权限生成新的解析树，寓意检查（字段ID是否存在）等

## 优化器

将前面的解析树转换成执行计划
分析使用哪个索引，如果有join一会分析使用哪个表作为驱动表等。

## 执行器

开始执行sql，首先会验证与没有相关权限。获取锁，打开表，通过meta数据，获取数据
# 存储引擎层

负责数据的存储与提取。其架构模式是插件式的，支持 `innodb`、`myisam`、`momory` 等。



# Mysql 查询

## 书写顺序
```sql
select distinct * from '表名' where '限制条件'  group by '分组依据' having '过滤条件' order by  limit '展示条数'
```

## 执行顺序

- from       -- 查询
- where      -- 限制条件
- group by   -- 分组
- having     -- 过滤条件
- order by   -- 排序
- limit      -- 展示条数
- distinct   -- 去重
- select     -- 查询的结果

## `LIMIT`

一下语句相同，跳过3行数据，查询10行数据
```sql
select * from  user where uid > 100 limit 10 offset 3;
select * from user where uid > 100 limit 3,10;
```

## `DISTINCT`

- `DISTINCT` 必须放在所有列名最前面
- `DISTCNCT` 作用于所有列而不是其后面的一个


## 正则

```sql
select * from emp where name regexp '^j.*(n|y)$';
```

## 使用通配符 `*`

- 除非你确实需要表中的每一列否则检索不需要的列通常会降低检索和应用程序的性能
- 网络IO：内存、CPU、磁盘的开销基本是微秒的影响很小但是网络IO可能带来秒级延迟
- 索引问题

```sql
select age from user where uid = 100;
select * from user where uid = 100;
```

如果`age`字段有索引，第一个sql直接使用索引中的值就是结果,第二个sql需要二次回表查询，造成额外性能开销。


## `ORDER BY`

### `ORDER BY` 排序使用的列可以是非检索字段,如下`uid`字段不在`select`查询字段中

```sql
select age from user order by uid;
```

### 多字段排序
```sql
select age, nickname from user order by uid,age;
```

### 相对位置排序
```sql
select age, nickname from user order by 1,2;
```
- 相当于按 `age`,`nickname` 排序
- 注意当排序字段不在`SELECT` 申明的字段后面不可以使用

### 指定排序方向

```sql
select age, nickname from user order by age, nickname asc;
select age, nickname from user order by age, nickname desc;
```





## 集合查询

`max` 、`min` 、`avg` 、`sum` 、`count` 、`group_concat` 

## 联表查询

- 内连接：inner join
- 左连接：left join
- 右连接：right join
- 全连接： 左连接 union 右连接
- replace 替换

## 拼接

- concat
- concat_ws
- group_concat
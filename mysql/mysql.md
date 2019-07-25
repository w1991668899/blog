# 什么是事务?

ACID  
原子性： 事务执行过程中不可以被中断，要么成功，要么失败回滚
一致性： 事务前后对数据库没有影响
隔离性： 多个事务同时执行，不会相互干扰
持久性： 事务成功后数据持久化存储


# 事务隔离级别
读提交： 事务提交之后它做的变更才能被其他事务看到
读未提交： 事务还未提交他做的变更就能被其他事务看到
可重复度：一个事务执行过程中看到的数据总是同这个事务启动时看到的数据是一致的，同时未提交的事务数据变更对其他事务不可见
串行化：加锁顺序执行

# 并发事务带来哪些问题？

脏读：事务A在执行中将data进行了修改但事务A还未提交此时事务B读取了修改后还未提交的数据
不可重复度：事务A在执行中多次读取data，在这个过程中事务B也访问了data
幻读: 事务A在执行过程中读取了部分数据，然后事务插入了部分新数据，当事务A再次查询时就会发现多了一些数据
丢失修改：

# 左前缀原则

# 表修稿语句

alter table `table_name` add primary key(`column`);

alter table `table_name` add unique(`column`);

alter table `table_name` add index index_name(`column`);

alter table `table_name` add fulltext(`column`);

alter table `table_name` add index index_name(`column1`, `column2`, `column3`);















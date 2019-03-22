# linux 性能优化基本思路

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/linux/linux%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%E5%9F%BA%E6%9C%AC%E6%80%9D%E8%B7%AF.jpeg'>
</p>

平时都在linux或mac下进行开发，最近看了一些关于linux性能优化分析的知识，这里做些总结。

参考:<br>

- 《linux性能优化》
- 《高性能linux》

linux性能优化需要我们理解应用程序的与系统的一些基本原理，再结合工作中的一些实践，不断练习，建立起一个整体的结构。这时候它就不再难了，因为一个问题当你知道从哪里入手它就解决了一半了。

一般思路:<br>

应用程序、库函数、系统调用、内核、硬件<br><br>
CPU 、磁盘IO 、内存、网路

# sysstat性能监控工具包

这里介绍一些我常用的命令（ubuntu系统）

## mpstat输出所有CPU的平均统计信息
```
root@e6fb5de4bca2:/# mpstat
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:41:47     CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
09:41:47     all    0.07    0.00    0.26    0.00    0.00    0.00    0.00    0.00    0.00   99.67
root@e6fb5de4bca2:/#
```

## mpstat -P ALL  0开始独立的输出每个CPU的统计信息，0表示第一个cpu

```
root@e6fb5de4bca2:/# mpstat -P ALL
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:46:12     CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
09:46:12     all    0.07    0.00    0.26    0.00    0.00    0.00    0.00    0.00    0.00   99.67
09:46:12       0    0.06    0.00    0.16    0.00    0.00    0.00    0.00    0.00    0.00   99.78
09:46:12       1    0.11    0.00    0.18    0.01    0.00    0.00    0.00    0.00    0.00   99.71
09:46:12       2    0.09    0.00    0.45    0.00    0.00    0.00    0.00    0.00    0.00   99.46
09:46:12       3    0.04    0.00    0.14    0.00    0.00    0.00    0.00    0.00    0.00   99.81
09:46:12       4    0.03    0.00    0.13    0.00    0.00    0.00    0.00    0.00    0.00   99.83
09:46:12       5    0.10    0.00    0.47    0.00    0.00    0.00    0.00    0.00    0.00   99.43
root@e6fb5de4bca2:/#
```

##  mpstat -P ALL 3 同上，每隔3秒输出一次
```
root@e6fb5de4bca2:/# mpstat -P ALL 3
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:48:11     CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
09:48:14     all    0.17    0.00    0.11    0.00    0.00    0.00    0.00    0.00    0.00   99.72
09:48:14       0    0.00    0.00    0.34    0.00    0.00    0.00    0.00    0.00    0.00   99.66
09:48:14       1    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
09:48:14       2    0.33    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00   99.67
09:48:14       3    0.00    0.00    0.33    0.00    0.00    0.00    0.00    0.00    0.00   99.67
09:48:14       4    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00  100.00
09:48:14       5    0.67    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00   99.33
```

## pidstat 输出所有活跃的任务
```
root@e6fb5de4bca2:/# pidstat
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:51:36      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:51:36        0         1    0.00    0.00    0.00    0.00    0.00     4  bash
09:51:36        0        16    0.00    0.00    0.00    0.00    0.00     2  bash
09:51:36        0        59    0.00    0.00    0.00    0.00    0.00     0  bash
09:51:36        0        76    0.00    0.00    0.00    0.00    0.00     2  bash
09:51:36        0       116    0.00    0.00    0.00    0.00    0.00     2  bash
09:51:36        0       140    0.00    0.00    0.00    0.00    0.00     4  bash
09:51:36        0      4640    0.00    0.00    0.00    0.00    0.00     4  bash
09:51:36        0      4692    0.00    0.00    0.00    0.00    0.00     2  bash
root@e6fb5de4bca2:/#
```
## pidstat -p ALL  输出所有活跃和非活跃的任务
```
root@e6fb5de4bca2:/# pidstat -p ALL
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:53:43      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:53:43        0         1    0.00    0.00    0.00    0.00    0.00     4  bash
09:53:43        0        16    0.00    0.00    0.00    0.00    0.00     2  bash
09:53:43        0        59    0.00    0.00    0.00    0.00    0.00     0  bash
09:53:43        0        76    0.00    0.00    0.00    0.00    0.00     2  bash
09:53:43        0       116    0.00    0.00    0.00    0.00    0.00     2  bash
09:53:43        0       140    0.00    0.00    0.00    0.00    0.00     4  bash
09:53:43        0       155    0.00    0.00    0.00    0.00    0.00     1  bash
09:53:43        0      4640    0.00    0.00    0.00    0.00    0.00     4  bash
09:53:43        0      4692    0.00    0.00    0.00    0.00    0.00     2  bash
09:53:43        0      4736    0.00    0.00    0.00    0.00    0.00     4  pidstat
root@e6fb5de4bca2:/#
```

## 以秒为单位对IO统计信息进行刷新
```
root@e6fb5de4bca2:/# pidstat -d 2
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:55:28      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command

09:55:30      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command

09:55:32      UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
```

## pidstat -t -p 1 2 3   每隔2秒输出1号进程信息，输出3次
```
root@e6fb5de4bca2:/# pidstat -t -p 1 2 3
Linux 4.9.125-linuxkit (e6fb5de4bca2) 	03/22/19 	_x86_64_	(6 CPU)

09:57:55      UID      TGID       TID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:57:57        0         1         -    0.00    0.00    0.00    0.00    0.00     4  bash
09:57:57        0         -         1    0.00    0.00    0.00    0.00    0.00     4  |__bash

09:57:57      UID      TGID       TID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:57:59        0         1         -    0.00    0.00    0.00    0.00    0.00     4  bash
09:57:59        0         -         1    0.00    0.00    0.00    0.00    0.00     4  |__bash
```

**其他命令的使用请自行查询**


# stress压测工具

```
root@e6fb5de4bca2:/# apt install stress    // 安装stress
Reading package lists... Done
Building dependency tree
Reading state information... Done
stress is already the newest version (1.0.4-2).
0 upgraded, 0 newly installed, 0 to remove and 8 not upgraded.   // 已经安装过
root@e6fb5de4bca2:/# stress             // 查看是否安装成功
`stress' imposes certain types of compute stress on your system

Usage: stress [OPTION [ARG]] ...
 -?, --help         show this help statement
     --version      show version statement
 -v, --verbose      be verbose
 -q, --quiet        be quiet
 -n, --dry-run      show what would have been done
 -t, --timeout N    timeout after N seconds
     --backoff N    wait factor of N microseconds before work starts
 -c, --cpu N        spawn N workers spinning on sqrt()
 -i, --io N         spawn N workers spinning on sync()
 -m, --vm N         spawn N workers spinning on malloc()/free()
     --vm-bytes B   malloc B bytes per vm worker (default is 256MB)
     --vm-stride B  touch a byte every B bytes (default is 4096)
     --vm-hang N    sleep N secs before free (default none, 0 is inf)
     --vm-keep      redirty memory instead of freeing and reallocating
 -d, --hdd N        spawn N workers spinning on write()/unlink()
     --hdd-bytes B  write B bytes per hdd worker (default is 1GB)

Example: stress --cpu 8 --io 4 --vm 2 --vm-bytes 128M --timeout 10s

Note: Numbers may be suffixed with s,m,h,d,y (time) or B,K,M,G (size).
root@e6fb5de4bca2:/#
```

## stress 基本命令

**CPU测试**<br>
```
root@e6fb5de4bca2:/# stress -c 4            // 增加4个cpu进程，处理sqrt()函数函数，以提高系统CPU负荷
stress: info: [4666] dispatching hogs: 4 cpu, 0 io, 0 vm, 0 hdd
```
**内存测试**<br>
```
root@e6fb5de4bca2:/# stress -i 4 --vm 10 --vm-bytes 1G --vm-hang 100 --timeout 100s   // 新增4个io进程，10个内存分配进程，每次分配大小1G，分配后不释放，测试100S
stress: info: [4674] dispatching hogs: 0 cpu, 4 io, 10 vm, 0 hdd
```

**磁盘IO测试**<br>
```
root@e6fb5de4bca2:/# stress -d 1 --hdd-bytes 3G             // 新增1个写进程，每次写3G文件块
stress: info: [4689] dispatching hogs: 0 cpu, 0 io, 0 vm, 1 hdd
```


# 平均负载

平均负载我的理解就是系统处于可运行状态与不可中断状态的平均进程数量。

```
wt-001% uptime 
 17:02:55 up 65 days, 22:52,  2 users,  load average: 0.15, 0.24, 0.19
wt-001% 

// 当前时间 登陆时间 正在登陆用户数 过去1分钟，5分钟，15分钟的平均负载
```

# 平均负载与CPU使用率的区别

CPU密集型进程会使CPU平均负载升高，IO密集型也会使平均负载升高，但是CPU使用率不一定升高，要注意区别，不要混淆。
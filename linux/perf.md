# Perf

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/linux/perf.jpeg'>
</p>

## 安装与使用的问题

```
wt-001% sudo apt install linux-tools   // 开始安装
[sudo] password for wt:
Reading package lists... Done
Building dependency tree
Reading state information... Done
Package linux-tools is a virtual package provided by:
  linux-tools-virtual-hwe-18.04-edge 4.18.0.13.62
  linux-tools-virtual-hwe-18.04 4.18.0.13.63
  linux-tools-virtual 4.15.0.43.45
  linux-tools-oem 4.15.0.1030.35
  linux-tools-lowlatency-hwe-18.04-edge 4.18.0.13.62
  linux-tools-lowlatency-hwe-18.04 4.18.0.13.63
  linux-tools-lowlatency 4.15.0.43.45
  linux-tools-gke 4.15.0.1026.28
  linux-tools-generic-hwe-18.04-edge 4.18.0.13.62
  linux-tools-generic-hwe-18.04 4.18.0.13.63
  linux-tools-generic 4.15.0.43.45
  linux-tools-gcp 4.15.0.1026.28
  linux-tools-aws 4.15.0.1031.30
You should explicitly select one to install.

E: Package 'linux-tools' has no installation candidate
wt-001%

wt-001% sudo apt install linux-tools-virtual-hwe-18.04-edge 4.18.0.13.62  // 选择正确内核版本安装
Reading package lists... Done
Building dependency tree
Reading state information... Done
E: Unable to locate package 4.18.0.13.62
E: Couldn't find any package by glob '4.18.0.13.62'
E: Couldn't find any package by regex '4.18.0.13.62'
wt-001%
```

我在Ubuntu系统中安装perf碰到了上面的问题，
# net包源码解析 `ip.go` 文件

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/net_01.jpeg'>
</p>

# 常量

在 `ip.go` 文件中
```
// IP address lengths (bytes).
const (
	IPv4len = 4         
	IPv6len = 16
)
```

# 常见变量

## 常用ipv4地址

```
// Well-known IPv4 addresses
var (
	IPv4bcast     = IPv4(255, 255, 255, 255) // limited broadcast 广播地址
	IPv4allsys    = IPv4(224, 0, 0, 1)       // all systems 所有系统，包括主机和路由器，这是一个组播地址
	IPv4allrouter = IPv4(224, 0, 0, 2)       // all routers 所有组播路由器
	IPv4zero      = IPv4(0, 0, 0, 0)         // all zeros 本地网络，只能作为本地源地址其才是合法的
)
```

## 常用ipv6地址

```
// Well-known IPv6 addresses
var (
	IPv6zero                   = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6unspecified            = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	IPv6loopback               = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallnodes      = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	IPv6linklocalallrouters    = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```


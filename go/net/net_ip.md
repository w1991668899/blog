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

## IP知识讲解 [wiki](https://zh.wikipedia.org/wiki/IPv4)
 
最初，一个IP地址被分成两部分：网络识别码在地址的高位字节中，主机识别码在剩下的部分中。
<br><br>为了克服这个限制，在随后出现的分类网络中，地址的高位字节被重定义为网络的类(Class)。这个系统定义了五个类别：A、B、C、D和E。A、B和C类有不同的网络类别长度，剩余的部分被用来识别网络内的主机，这就意味着每个网络类别有着不同的给主机编址的能力。D类被用于多播地址，E类被留作将来使用。
<br><br>1993年，无类别域间路由（CIDR）正式地取代了分类网络，后者也因此被称为“有类别”的。
<br><br>CIDR被设计为可以重新划分地址空间，因此小的或大的地址块均可以分配给用户。CIDR创建的分层架构由互联网号码分配局（IANA）和区域互联网注册管理机构（RIR）进行管理，每个RIR均维护着一个公共的WHOIS数据库，以此提供IP地址分配的详情。

## 无类型域间选路 CIDR

这种方式打破了原来设计的几类地址分类，将32位IP地址一分为二，前面是网络号后面是主机号。
<br><br>如： 10.100.122.2/24 这种地址形式就是CIDR。前24位是网络号，后8位是主机号。
<br><br>伴随CIDR出现的是一个**广播地址**，10.100.122.255 。如果发送这个地址，所有 10.100.122 网络里面的机器都可以收到。另一个是**子网掩码** 255.255.255.0
<br><br>**将子网掩码和IP地址按位计算AND，即可得到网络号**

```
// An IP is a single IP address, a slice of bytes.
// Functions in this package accept either 4-byte (IPv4)
// or 16-byte (IPv6) slices as input.
//
// Note that in this documentation, referring to an
// IP address as an IPv4 address or an IPv6 address
// is a semantic property of the address, not just the
// length of the byte slice: a 16-byte slice can still
// be an IPv4 address.
type IP []byte

// IP表示一个简单的IP地址，它是一个byte类型的slice，能够接受4字节（IPV4）或者16字节（IPV6）输入。
// 注意，IP地址是IPv4地址还是IPv6地址是语义上的特性，而不取决于切片的长度：16字节的切片也可以是IPv4地址。
```

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


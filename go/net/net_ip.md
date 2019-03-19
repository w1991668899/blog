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

// 很容易看出这表示ip地址的长度（bytes），其中ipv4长度是4，ipv6地址长度是16
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

## 公有IP地址与私有IP地址

如果你搭建一个网站使所有人都能访问那么就需要一个公有IP，公有IP是有组织统一分配的。
192.168.0.xxx 是最常用的私有IP，家里的WIFI路由器一般就是 192.168.0.1,而 192.168.0.255 就是广播地址。一旦发送这个地址整个 192.168.0网络中的机器都能收到。

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
// 如果是IPv4会自定义前面12个字节可查找源码
var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}
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

# 自定义类型



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

// An IP mask is an IP address.
type IPMask []byte

// An IPNet represents an IP network.
type IPNet struct { 
	IP   IP     // network number  网路地址
	Mask IPMask // network mask    子网掩码
}
```

```
// IPv4 returns the IP address (in 16-byte form) of the
// IPv4 address a.b.c.d.
func IPv4(a, b, c, d byte) IP {
	p := make(IP, IPv6len)
	copy(p, v4InV6Prefix)
	p[12] = a
	p[13] = b
	p[14] = c
	p[15] = d
	return p
}
// 获取ipv4地址

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// IPv4Mask returns the IP mask (in 4-byte form) of the
// IPv4 mask a.b.c.d.
func IPv4Mask(a, b, c, d byte) IPMask {
	p := make(IPMask, IPv4len)
	p[0] = a
	p[1] = b
	p[2] = c
	p[3] = d
	return p
}
// 返回ip掩码,其中ip掩码形式是ipv4掩码(4 byte模式)a.b.c.d

// CIDRMask returns an IPMask consisting of `ones' 1 bits
// followed by 0s up to a total length of `bits' bits.
// For a mask of this form, CIDRMask is the inverse of IPMask.Size.
func CIDRMask(ones, bits int) IPMask {
	if bits != 8*IPv4len && bits != 8*IPv6len {
		return nil
	}
	if ones < 0 || ones > bits {
		return nil
	}
	l := bits / 8
	m := make(IPMask, l)
	n := uint(ones)
	for i := 0; i < l; i++ {
		if n >= 8 {
			m[i] = 0xff
			n -= 8
			continue
		}
		m[i] = ^byte(0xff >> n)
		n = 0
	}
	return m
}
返回一个CIDRMask,其中CIDRMask总bit数目是bits,前ones位是1,其余位是0.
```

看代码
```
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv4(8, 8, 9, 20).String())
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))
	fmt.Println(net.IPv4Mask(255, 255, 255, 0).Size())
	fmt.Println(net.IPv4Mask(255, 255, 255, 0).String())
	fmt.Println(net.CIDRMask(31, 32))
	fmt.Println(net.CIDRMask(31, 32).Size())
	fmt.Println(net.CIDRMask(64, 128).String())
}

//返回值
8.8.9.20
ffffff00
24 32
ffffff00
fffffffe
31 32
ffffffffffffffff0000000000000000

```

```
// IsUnspecified reports whether ip is an unspecified address, either
// the IPv4 address "0.0.0.0" or the IPv6 address "::".
func (ip IP) IsUnspecified() bool {
	return ip.Equal(IPv4zero) || ip.Equal(IPv6unspecified)
}
// 报告ip是否是未指定的地址，IPv4地址“0.0.0.0”或IPv6地址“::”。

// IsLoopback reports whether ip is a loopback address.
func (ip IP) IsLoopback() bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 127
	}
	return ip.Equal(IPv6loopback)
}
// IsLoopback报告ip是否是环回地址。

// IsMulticast reports whether ip is a multicast address.
func (ip IP) IsMulticast() bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0]&0xf0 == 0xe0
	}
	return len(ip) == IPv6len && ip[0] == 0xff
}
// 报告ip是否是多播地址

// IsInterfaceLocalMulticast reports whether ip is
// an interface-local multicast address.
func (ip IP) IsInterfaceLocalMulticast() bool {
	return len(ip) == IPv6len && ip[0] == 0xff && ip[1]&0x0f == 0x01
}
// 报告ip是否是接口本地多播地址。

// IsLinkLocalMulticast reports whether ip is a link-local
// multicast address.
func (ip IP) IsLinkLocalMulticast() bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 224 && ip4[1] == 0 && ip4[2] == 0
	}
	return len(ip) == IPv6len && ip[0] == 0xff && ip[1]&0x0f == 0x02
}
// 报告ip是否是链路本地多播地址。

// IsLinkLocalUnicast reports whether ip is a link-local
// unicast address.
func (ip IP) IsLinkLocalUnicast() bool {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 169 && ip4[1] == 254
	}
	return len(ip) == IPv6len && ip[0] == 0xfe && ip[1]&0xc0 == 0x80
}
// 报告ip是否是链路本地单播地址。

// IsGlobalUnicast reports whether ip is a global unicast
// address.
//
// The identification of global unicast addresses uses address type
// identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
// the exception of IPv4 directed broadcast addresses.
// It returns true even if ip is in IPv4 private address space or
// local IPv6 unicast address space.
func (ip IP) IsGlobalUnicast() bool {
	return (len(ip) == IPv4len || len(ip) == IPv6len) &&
		!ip.Equal(IPv4bcast) &&
		!ip.IsUnspecified() &&
		!ip.IsLoopback() &&
		!ip.IsMulticast() &&
		!ip.IsLinkLocalUnicast()
}
// 报告ip是否是全球单播地址。
// 全局单播地址的标识使用RFC 1122，RFC 4632和RFC 4291中定义的地址类型标识，但IPv4定向广播地址除外。即使ip位于IPv4专用地址空间或本地IPv6单播地址空间，它也会返回true。

// To4 converts the IPv4 address ip to a 4-byte representation.
// If ip is not an IPv4 address, To4 returns nil.
func (ip IP) To4() IP {
	if len(ip) == IPv4len {
		return ip
	}
	if len(ip) == IPv6len &&
		isZeros(ip[0:10]) &&
		ip[10] == 0xff &&
		ip[11] == 0xff {
		return ip[12:16]
	}
	return nil
}
// To4将IPv4地址ip转换为4字节表示形式。如果ip不是IPv4地址，则To4返回nil。

// To16 converts the IP address ip to a 16-byte representation.
// If ip is not an IP address (it is the wrong length), To16 returns nil.
func (ip IP) To16() IP {
	if len(ip) == IPv4len {
		return IPv4(ip[0], ip[1], ip[2], ip[3])
	}
	if len(ip) == IPv6len {
		return ip
	}
	return nil
}
// 将IP地址ip转换为16字节的表示形式。如果ip不是IP地址（它是错误的长度），To16返回nil。

// DefaultMask returns the default IP mask for the IP address ip.
// Only IPv4 addresses have default masks; DefaultMask returns
// nil if ip is not a valid IPv4 address.
func (ip IP) DefaultMask() IPMask {
	if ip = ip.To4(); ip == nil {
		return nil
	}
	switch {
	case ip[0] < 0x80:
		return classAMask
	case ip[0] < 0xC0:
		return classBMask
	default:
		return classCMask
	}
}
// 返回IP地址ip的默认IP掩码。只有IPv4地址具有默认掩码; 如果ip不是有效的IPv4地址，则DefaultMask返回nil。

// Mask returns the result of masking the IP address ip with mask.
func (ip IP) Mask(mask IPMask) IP {
	if len(mask) == IPv6len && len(ip) == IPv4len && allFF(mask[:12]) {
		mask = mask[12:]
	}
	if len(mask) == IPv4len && len(ip) == IPv6len && bytealg.Equal(ip[:12], v4InV6Prefix) {
		ip = ip[12:]
	}
	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}
// 返回用掩码掩码IP地址ip的结果。

// String returns the string form of the IP address ip.
// It returns one of 4 forms:
//   - "<nil>", if ip has length 0
//   - dotted decimal ("192.0.2.1"), if ip is an IPv4 or IP4-mapped IPv6 address
//   - IPv6 ("2001:db8::1"), if ip is a valid IPv6 address
//   - the hexadecimal form of ip, without punctuation, if no other cases apply
func (ip IP) String() string {
	p := ip

	if len(ip) == 0 {
		return "<nil>"
	}

	// If IPv4, use dotted notation.
	if p4 := p.To4(); len(p4) == IPv4len {
		const maxIPv4StringLen = len("255.255.255.255")
		b := make([]byte, maxIPv4StringLen)

		n := ubtoa(b, 0, p4[0])
		b[n] = '.'
		n++

		n += ubtoa(b, n, p4[1])
		b[n] = '.'
		n++

		n += ubtoa(b, n, p4[2])
		b[n] = '.'
		n++

		n += ubtoa(b, n, p4[3])
		return string(b[:n])
	}
	if len(p) != IPv6len {
		return "?" + hexString(ip)
	}

	// Find longest run of zeros.
	e0 := -1
	e1 := -1
	for i := 0; i < IPv6len; i += 2 {
		j := i
		for j < IPv6len && p[j] == 0 && p[j+1] == 0 {
			j += 2
		}
		if j > i && j-i > e1-e0 {
			e0 = i
			e1 = j
			i = j
		}
	}
	// The symbol "::" MUST NOT be used to shorten just one 16 bit 0 field.
	if e1-e0 <= 2 {
		e0 = -1
		e1 = -1
	}

	const maxLen = len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	b := make([]byte, 0, maxLen)

	// Print with possible :: in place of run of zeros
	for i := 0; i < IPv6len; i += 2 {
		if i == e0 {
			b = append(b, ':', ':')
			i = e1
			if i >= IPv6len {
				break
			}
		} else if i > 0 {
			b = append(b, ':')
		}
		b = appendHex(b, (uint32(p[i])<<8)|uint32(p[i+1]))
	}
	return string(b)
}
// 返回IP地址ip的字符串形式

// MarshalText implements the encoding.TextMarshaler interface.
// The encoding is the same as returned by String, with one exception:
// When len(ip) is zero, it returns an empty slice.
func (ip IP) MarshalText() ([]byte, error) {
	if len(ip) == 0 {
		return []byte(""), nil
	}
	if len(ip) != IPv4len && len(ip) != IPv6len {
		return nil, &AddrError{Err: "invalid IP address", Addr: hexString(ip)}
	}
	return []byte(ip.String()), nil
}
// 实现了encoding.TextMarshaler接口。编码与String返回的一样，但有一个例外：当len（ip）为零时，它返回一个空片。

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The IP address is expected in a form accepted by ParseIP.
func (ip *IP) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*ip = nil
		return nil
	}
	s := string(text)
	x := ParseIP(s)
	if x == nil {
		return &ParseError{Type: "IP address", Text: s}
	}
	*ip = x
	return nil
}
// 实现了encoding.TextUnmarshaler接口。IP地址预计采用ParseIP接受的形式。

// Equal reports whether ip and x are the same IP address.
// An IPv4 address and that same address in IPv6 form are
// considered to be equal.
func (ip IP) Equal(x IP) bool {
	if len(ip) == len(x) {
		return bytealg.Equal(ip, x)
	}
	if len(ip) == IPv4len && len(x) == IPv6len {
		return bytealg.Equal(x[0:12], v4InV6Prefix) && bytealg.Equal(ip, x[12:])
	}
	if len(ip) == IPv6len && len(x) == IPv4len {
		return bytealg.Equal(ip[0:12], v4InV6Prefix) && bytealg.Equal(ip[12:], x)
	}
	return false
}
// 报告ip和x是否是相同的IP地址。IPv4地址和IPv6形式的相同地址被认为是相同的。

```


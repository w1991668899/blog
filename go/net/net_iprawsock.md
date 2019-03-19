# net包源码解析 `iprawsock.go` `net.go`文件

<p align='center'>
<img src='https://github.com/w1991668899/blog/blob/master/image/go/net_iprawsock.jpg'>
</p>

# IPAddr 类型

```
// BUG(mikio): On every POSIX platform, reads from the "ip4" network
// using the ReadFrom or ReadFromIP method might not return a complete
// IPv4 packet, including its header, even if there is space
// available. This can occur even in cases where Read or ReadMsgIP
// could return a complete packet. For this reason, it is recommended
// that you do not use these methods if it is important to receive a
// full packet.
//
// The Go 1 compatibility guidelines make it impossible for us to
// change the behavior of these methods; use Read or ReadMsgIP
// instead.

// BUG(mikio): On JS, NaCl and Plan 9, methods and functions related
// to IPConn are not implemented.

// BUG(mikio): On Windows, the File method of IPConn is not
// implemented.

// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP   IP
	Zone string // IPv6 scoped addressing zone
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string { return "ip" }

func (a *IPAddr) String() string {
	if a == nil {
		return "<nil>"
	}
	ip := ipEmptyString(a.IP)
	if a.Zone != "" {
		return ip + "%" + a.Zone
	}
	return ip
}

func (a *IPAddr) isWildcard() bool {
	if a == nil || a.IP == nil {
		return true
	}
	return a.IP.IsUnspecified()
}

func (a *IPAddr) opAddr() Addr {
	if a == nil {
		return nil
	}
	return a
}

// ResolveIPAddr returns an address of IP end point.
//
// The network must be an IP network name.
//
// If the host in the address parameter is not a literal IP address,
// ResolveIPAddr resolves the address to an address of IP end point.
// Otherwise, it parses the address as a literal IP address.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveIPAddr(network, address string) (*IPAddr, error) {
	if network == "" { // a hint wildcard for Go 1.0 undocumented behavior
		network = "ip"
	}
	afnet, _, err := parseNetwork(context.Background(), network, false)
	if err != nil {
		return nil, err
	}
	switch afnet {
	case "ip", "ip4", "ip6":
	default:
		return nil, UnknownNetworkError(network)
	}
	addrs, err := DefaultResolver.internetAddrList(context.Background(), afnet, address)
	if err != nil {
		return nil, err
	}
	return addrs.forResolve(network, address).(*IPAddr), nil
}
// 将ip地址解析成形如"host"或者"ipv6-host%zone"的地址形式,解析域名必须在指定的网络中,指定网络包括ip,ip4或者ip6
// 根据指定的 network ip协议对地址进行解析
```
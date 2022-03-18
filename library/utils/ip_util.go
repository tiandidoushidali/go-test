package utils

import "net"

type ipNetType uint8
const (
	IP_LOOKBACK ipNetType = 0 + iota
	IP_PRIVATE
	IP_PUBLIC
)


// 检测ip类型
func DetectIpNetType(ip net.IP) ipNetType {
	if ip.IsLoopback() {
		return IP_LOOKBACK
	}
	var (
		// rfc1918
		rfc1918 = []*net.IPNet{
			{IP: net.IPv4(0x0A, 0x00, 0x00, 0x00), Mask: net.CIDRMask(0x08, 32)}, // 10.0.0.0/8
			{IP: net.IPv4(0xAC, 0x10, 0x00, 0x00), Mask: net.CIDRMask(0x0C, 32)}, // 172.16.0.0/12
			{IP: net.IPv4(0xC0, 0xA8, 0x00, 0x00), Mask: net.CIDRMask(0x10, 32)}, // 192.168.0.0/16
		}
	)
	for _, ipNet := range rfc1918 {
		if ipNet.Contains(ip) {
			return IP_PRIVATE
		}
	}
	return IP_PUBLIC
}
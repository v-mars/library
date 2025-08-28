package utils

import (
	"fmt"
	"net"
	"strings"
)

// IPInSubnet ; 支持IPv4和IPv6
// ipAddr: 192.168.0.1;
// cidr: 192.168.0.0/24;
func IPInSubnet(ipAddr, cidr string) (bool, error) {
	ip := net.ParseIP(ipAddr)
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, err
	}
	if ipNet.Contains(ip) {
		return true, nil
	}
	return false, nil
}
func IPsInSubnets(ips, cidrs []string) (bool, error) {
	for _, ip := range ips {
		for _, cidr := range cidrs {
			inSubnet, err := IPInSubnet(ip, cidr)
			if err != nil {
				return false, err
			}
			if inSubnet {
				return true, nil
			}
		}
	}
	return false, nil
}

// ParseHost
func ParseHost(host string) (ip string, port string, err error) {
	// 如果host不包含端口，添加一个默认端口以便解析
	if !strings.Contains(host, ":") {
		host = net.JoinHostPort(host, "22") // 使用默认端口80
	}

	// 使用net.SplitHostPort解析host和端口
	ip, port, err = net.SplitHostPort(host)
	if err != nil {
		return "", "", err
	}

	return ip, port, nil
}

// ParseHostWithResolve 或者使用net.ResolveTCPAddr方法
func ParseHostWithResolve(host string) (ip string, port string, err error) {
	// 如果host不包含端口，添加默认端口
	if !strings.Contains(host, ":") {
		host = host + ":22"
	}

	// 解析TCP地址
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return "", "", err
	}

	return addr.IP.String(), fmt.Sprintf("%d", addr.Port), nil
}

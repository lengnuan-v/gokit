// +----------------------------------------------------------------------
// | host方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import (
	"encoding/binary"
	"net"
	"os"
	"runtime"
)

// 获取当前操作系统
func Goos() string {
	return runtime.GOOS
}

// 获取主机名
func Gethostname() (string, error) {
	return os.Hostname()
}

// 返回主机名对应的 IPv4地址
// Gethostbyname("localhost")
// Gethostbyname('www.example.com')
func Gethostbyname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// 将长整数地址转换为（IPv4）Internet标准点分格式的字符串
// IP2long("8.8.8.8")
func IP2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// 将长整数地址转换为（IPv4）Internet标准点分格式的字符串
// Long2ip(134744072)
func Long2ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

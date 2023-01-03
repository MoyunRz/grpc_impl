package heartbeat

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type ThisHostInfo struct {
	IpAddr    string `json:"ip_addr"`
	Domain    string `json:"domain"`
	Port      string `json:"port"`
	Rule      string `json:"rule"`
	Prefix    string `json:"prefix"`
	TimeStamp int64  `json:"timestamp"`
}

var PubLocalHostInfo ThisHostInfo

func InitLocalHostInfo(ipAddr, domain, port, rule, prefix string) ThisHostInfo {
	PubLocalHostInfo.IpAddr = ipAddr
	PubLocalHostInfo.Domain = domain
	PubLocalHostInfo.Port = port
	PubLocalHostInfo.Rule = rule
	PubLocalHostInfo.Prefix = prefix
	PubLocalHostInfo.TimeStamp = time.Now().Unix()
	return PubLocalHostInfo
}

// GetOutBoundIP
// 获取本机地址
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func GetHostInfo() ThisHostInfo {

	var thinfo ThisHostInfo
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
		return thinfo
	}
	thinfo.IpAddr = ip
	thinfo.Domain = ip
	thinfo.Port = ""
	thinfo.TimeStamp = time.Now().Unix()
	return thinfo
}

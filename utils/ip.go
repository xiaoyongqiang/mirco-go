package utils

import (
	"bytes"
	"net"
	"net/http"
	"strings"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

func inIPRange(r ipRange, ipAddress net.IP) bool {
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}

	return false
}

var privateIPRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

func isPrivateSubnet(ipAddress net.IP) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateIPRanges {
			if inIPRange(r, ipAddress) {
				return true
			}
		}
	}

	return false
}

// GetIPAddress return the client ip
func GetIPAddress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-IP"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				continue
			} else {
				return ip
			}
		}
	}

	remoteAddr := strings.Split(r.RemoteAddr, ":")[0]
	realIP := net.ParseIP(remoteAddr)
	if realIP.IsGlobalUnicast() && !isPrivateSubnet(realIP) {
		return remoteAddr
	}

	return ""
}

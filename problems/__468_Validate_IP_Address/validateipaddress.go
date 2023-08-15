/*
Package __468_Validate_IP_Address
@Time : 2022/5/29 9:32
@Author : sunxy
@File : validateipaddress
@description:
*/
package __468_Validate_IP_Address

import (
	"strconv"
	"strings"
)

const (
	ipv4    = "IPv4"
	ipv6    = "IPv6"
	neither = "Neither"
)

func validIPAddress(queryIP string) string {
	// 3个 . 或者 7个 :
	subadds := strings.Split(queryIP, ".")
	// validate ipv4
	if len(subadds) == 4 {
		return validateIpv4(subadds)
	}
	subadds = strings.Split(queryIP, ":")
	// validate ipv6
	if len(subadds) == 8 {
		return validateIpv6(subadds)
	}
	return neither

}

func validateIpv4(subAdds []string) string {
	// 每个字符串转int
	// 大小 0 -- 255
	// 如果是0, len == 1
	for _, s := range subAdds {
		n, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return neither
		}
		if n < 0 || n > 255 {
			return neither
		}
		if n == 0 && len(s) != 1 {
			return neither
		}
		if n != 0 && []byte(s)[0] == '0' {
			return neither
		}
	}
	return ipv4
}

func validateIpv6(subAdds []string) string {
	// 按照:分割
	// 转16进制数字, 0 -- ffff
	// 每个的len 不能超过4
	for _, s := range subAdds {
		if len(s) > 4 || len(s) == 0 {
			return neither
		}
		n, err := strconv.ParseInt(s, 16, 0)
		if err != nil {
			return neither
		}
		if n < 0 || n > 0xffff {
			return neither
		}
	}
	return ipv6
}

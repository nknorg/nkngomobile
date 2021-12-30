package nkngomobile

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sort"
)

var ErrInvalidIP = errors.New("invalid IP address")

func IpToUint32(s string) (uint32, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0, ErrInvalidIP
	}

	ip = ip.To4()
	if ip == nil {
		return 0, ErrInvalidIP
	}

	return uint32(ip[3]) | uint32(ip[2])<<8 | uint32(ip[1])<<16 | uint32(ip[0])<<24, nil
}

func Uint32ToIP(n uint32) net.IP {
	return net.IPv4(byte(n>>24), byte(n>>16&0xFF), byte(n>>8)&0xFF, byte(n&0xFF))
}

// IpRangeToCIDR both side inclusive
func IpRangeToCIDR(start, end uint32) []string {
	if start > end {
		return nil
	}

	// use uint64 to prevent overflow
	ip := int64(start)
	tail := int64(0)
	cidr := make([]string, 0)

	// decrease mask bit
	for {
		// count number of tailing zero bits
		for ; tail < 32; tail++ {
			if (ip>>(tail+1))<<(tail+1) != ip {
				break
			}
		}
		if ip+(1<<tail)-1 > int64(end) {
			break
		}
		cidr = append(cidr, fmt.Sprintf("%s/%d", Uint32ToIP(uint32(ip)).String(), 32-tail))
		ip += 1 << tail
	}

	// increase mask bit
	for {
		for ; tail >= 0; tail-- {
			if ip+(1<<tail)-1 <= int64(end) {
				break
			}
		}
		if tail < 0 {
			break
		}
		cidr = append(cidr, fmt.Sprintf("%s/%d", Uint32ToIP(uint32(ip)).String(), 32-tail))
		ip += 1 << tail
		if ip-1 == int64(end) {
			break
		}
	}

	return cidr
}

func ExcludeRoute(ipArray *StringArray) *StringArray {
	ips := make([]uint32, 0, ipArray.Len())
	for _, ip := range ipArray.Elems() {
		ipnum, err := IpToUint32(ip)
		if err != nil {
			log.Fatal(err)
		}
		ips = append(ips, ipnum)
	}

	sort.Slice(ips, func(i, j int) bool { return ips[i] < ips[j] })

	min := uint32(0)
	max := uint32(4294967295)

	res := NewStringArray()
	for _, ip := range ips {
		cidr := IpRangeToCIDR(min, ip-1)
		for _, s := range cidr {
			res.Append(s)
		}
		min = ip + 1
	}
	cidr := IpRangeToCIDR(min, max)
	for _, s := range cidr {
		res.Append(s)
	}
	return res
}

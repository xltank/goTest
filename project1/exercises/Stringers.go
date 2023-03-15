package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (ip IPAddr) String() string {
	strArray := [4]string{}
	for i, v := range ip {
		strArray[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strArray[:], ",")
}

package main

import "fmt"

type IPAddr [4]byte

func (addr IPAddr)String() string{
	return fmt.Sprintf("%v.%v.%v.%v",int(addr[0]),int(addr[1]),int(addr[2]),int(addr[3]))
}
// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

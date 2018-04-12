package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) != 2 {
		// TODO Fprintf
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	//TODO  会把一个IPv4或者IPv6的地址转换成IP类型
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

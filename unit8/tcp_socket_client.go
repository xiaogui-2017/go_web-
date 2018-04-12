package main

import (
	"fmt"
	"os"
	"net"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	//TODO bug is here
	//tcpAddr, err := net.ResolveIPAddr("tcp4", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		//fmt.Println(os.Stderr, "Fatal error :%s ", err.Error())
		fmt.Fprintf(os.Stderr, "Fatal error :%s ", err.Error())
		os.Exit(1)
	}
}

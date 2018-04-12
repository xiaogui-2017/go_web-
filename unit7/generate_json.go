package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	//Serverslice继承Servers
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{"shanghai_vpn", "127.0.0.1"})
	s.Servers = append(s.Servers, Server{"beijing_vpn", "127.0.0.2"})
	//把一个结构直接json化
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err", err)
	}

	fmt.Println(string(b))
}

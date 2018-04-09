package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数， 默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	//r.form的本质是 type Values map[string][]string    区分大小写的case-sensitive
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello zhangzhiyuan")  // 这个是写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)  // 设置访问的路由
	err := http.ListenAndServe(":9090", nil)  // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

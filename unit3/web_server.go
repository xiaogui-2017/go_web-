package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数， 默认是不会解析的

	//r.form的本质是 type Values map[string][]string    区分大小写的case-sensitive
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	//r.form的本质是 type Values map[string][]string    区分大小写的case-sensitive
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello zhangzhiyuan") // 这个是写入到w的是输出到客户端的
}

func main() {
	//我们调用了HandlerFunc(f), 类似强制类型转换f成为HandlerFunc类型， 这样f就拥有了ServerHTTP方法
	http.HandleFunc("/", sayhelloName) // 设置访问的路由

	//handler就用以配置外部路由器的， 即外部路由器只要实现Handler接口就可以，我们可以在自己实现的路由器ServHTTP里面实现
	//简易的自定义路由功能
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

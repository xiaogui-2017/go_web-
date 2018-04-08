// goroutine是通过go语言runtime管理的一个线程管理器，通过go关键字实现了
// 单线程并发
// 多线程并行
package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// 让cpu把时间片让给别人，下次某个时候恢复执行该goroutine
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("hello") // 开一个新的Goroutines
	// goroutines运行在同一个进程里面， 共享内存数据，原则：不要通过共享来通信， 要通过通信来共享
	say("hello") // 当前goroutines执行
}

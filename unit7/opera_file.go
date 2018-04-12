/*
操作文件
 */
package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	//userFile := "zzy.text"
	//fout, err := os.Create(userFile)
	fout, err := os.Create("zzy.text")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("just a text\n")
		fout.Write([]byte("Just a test\n"))
	}

	file_read, err := os.Open("zzy.text")
	defer file_read.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// TODO  make用于内建函数的内存分配，map slice channel,并且返回一个有初始值的T类型
	buf := make([]byte, 1024)
	for {
		n, _ := file_read.Read(buf)
		if n == 0 {
			break
		}
		//TODO
		os.Stdout.Write(buf[:n])
	}

	time.Sleep(10 * time.Second)
	os.Remove("zzy.text")
}

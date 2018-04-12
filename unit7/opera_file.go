/*
操作文件
 */
package main

import (
	"os"
	"fmt"
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

}

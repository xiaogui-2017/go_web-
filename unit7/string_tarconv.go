/*
整数转换成字符串后， 添加到现有的字节数组中
 */

package main

import (
	"strconv"
	"fmt"
)

func main() {
	str := make([]byte, 0, 100)
	//将整数转换成字符串后， 添加到现有的字节数组中
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}

/*
format系列函数把其他类型的转换为字符串
 */
package main

import (
	"strconv"
	"fmt"
)

func main() {
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
}

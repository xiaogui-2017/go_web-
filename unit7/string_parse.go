/*
string  parse系列函数把字符串转换为其他类型
 */
package main

import (
	"strconv"
	"fmt"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("13245", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)

	fmt.Println(a, b, c, d, e)

}

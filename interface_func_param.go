//interface函数参数
package main

import (
	"strconv"
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return " " + h.name + " - " + strconv.Itoa(h.age) + "years - 电话" + h.phone + " "
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-xxx"}
	fmt.Println("This Human is : ", Bob)
}

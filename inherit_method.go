package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

// navigate to dumplicate  导航到重复
type Student struct {
	Human // 匿名函数
	school string
}

type Employee struct {
	Human // 匿名函数
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	//此处call为打电话， 而非调用
	fmt.Printf("Hi I'm %s, you can all me on %s \n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "111-111-111"}, "NWNU"}
	sam := Student{Human{"Sam", 45, "222-222-222"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}

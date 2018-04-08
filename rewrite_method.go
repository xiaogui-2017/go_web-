//method重写
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human
	company string
}

type Student struct {
	Human
	school string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s ,  Cal me on %s \n", h.name, h.phone)
}

//重写SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s , I workat %s, Cal me on %s \n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"zzy", 25, "1815363"}, "NWNU"}
	mark.SayHi()
	zzy := Employee{Human{"topsec", 20, "010-xxxx"}, "topsec"}
	zzy.SayHi()
}

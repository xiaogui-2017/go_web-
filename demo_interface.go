package main

import "fmt"

// 如果某个对象实现了某个接口的所有方法， 则此对象就实现了此接口。

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("la la la", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle ...", beerStein)
}

//重载Human的sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me %s\n", e.name, e.company, e.phone)
}

//student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // again and again and ...
}

// Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please !!! Get me through the day
}

//定义interface类型
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

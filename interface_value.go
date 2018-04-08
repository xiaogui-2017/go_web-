package main

import "fmt"

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

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me %s\n", h.name, h.phone)
}
func (h Human) Sing(lyrics string) {
	fmt.Println("la la la", lyrics)
}

//重载Human的sayhi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	// Student {Human{name, age, phone}, loan}
	mike := Student{Human{"Mike", 25, "222-222-222"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-222"}, "Harvard", 100}
	sam := Student{Human{"Sam", 36, "444-222-222"}, "Golang Inc", 1000}
	Tom := Student{Human{"Sam", 36, "222-444-222"}, "Things Ltd.", 5000}
	//定义Men类型变量i
	//i就是interface
	var i Men
	i = mike
	fmt.Println("this is Mike, a student:")
	i.SayHi()
	i.Sing("November rain")

	//i也存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee: ")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	// 接口和slice结合起来
	fmt.Println("Let's use a slice of Men and see what happens")
	//开辟一些内存出来
	x := make([]Men, 3)
	// T这三个都是不同的元素， 但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
	//interface是一组抽象方法的集合， 需要有其他类型非interface类型实现， 而不能自我实现
	// 当看见一只鸟走起来像鸭子， 游泳起来像鸭子， 叫起来也像鸭子， 那么这只鸟就可以被称为鸭子
}

package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 函数变成方法,方法可以方位接受者(结构体)的字段,就像结构体里访问字段一样
// area方法属于r和c
// method可以定义在任何自定义的类型、内置类型、结构体等等

// 可根据自己的需求使用值传递或者指针传递
func (r Rectangle) area() float64 {
	return r.width * r.height
}
// 可根据自己的需求使用值传递或者指针传递
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// TODO
func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human //匿名字段，默认Student包含Human的所有字段，相当于继承
	specially string
}

func main() {
	// 初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	//查询
	fmt.Println("His name is", mark.name)
	fmt.Println("His age is", mark.age)
	fmt.Println("His anme weight", mark.weight)
	fmt.Println("His anme specially", mark.specially)
	//修改
	fmt.Println("_____________________________________")
	mark.specially = "AI"
	fmt.Println("After His anme specially", mark.specially)
	mark.age = 38
	fmt.Println("After His age is", mark.age)
	mark.weight += 60
	fmt.Println("After His anme weight", mark.weight)
}

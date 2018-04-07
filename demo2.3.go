package main

import "fmt"

func add1(a int) int {
	a = a + 1 // 我们改变了a的值, 传值就是传入了一份拷贝
	return a
}

func main() {
	x := 3
	x1 := add1(x) // 操作的是x的拷贝，不改变原来的值, 传给函数，拷贝一份
	fmt.Println("x+1=", x1)
	fmt.Println("x = ", x)  // 输出原来的值
}

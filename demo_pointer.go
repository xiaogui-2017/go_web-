package main

import "fmt"

func main() {
	x := 1
	// &：为了操作自己，把自己的地址传给调用的函数，类型为指针
	// 有时候也可以传入结构体，因为不用指针的话，是传值操作，深拷贝， 复制大量的结构体太过消耗性能
	// slice(只有改变长度时候用&) string map机制类似指针所以传入指针的时候不需要取值符号&
	fmt.Println(Add(&x))
	fmt.Println(x)
}
// 星号的意思是，他为了操作自己，传递是地址，星号会找到地址的值（相当于解包）
// 如果不用地址的话，i就是地址
func Add(i *int) int {
	*i = *i + 1
	return *i
}

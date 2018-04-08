//自定义方法
package main

import (
	"fmt"
)

//自定义类型就是别名，Color就是byte的别名
type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}

//boxlist嵌套的是box
type BoxList []Box // a slice of boxes

//Volume定义接受者为Box
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

//然后以上面的自定义类型为接受者定义了一些method ，c就是字节模式, 亦可以用*Box（指针）作为receiver
func (b Box) SetColor(c Color) {
	b.color = c  // 等价于*b.color = c
}

//返回list中容量最大的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	//将WHITE强制转换为byte类型,并赋值给K
	k := Color(WHITE)
	for _, b := range bl {
		//找出最大的容量赋值给v, 对应的颜色赋值给k
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		// TODO 为什么会err, 参数本应为byte字节,uint8
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := [] string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	//返回byte编程字符串格式
	return strings[c]
}

func main() {
	boxes := BoxList{
		// TODO 此处所有的颜色应该为byte，即为uint8
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, WHITE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("we hava %d boxes in our set\n", len(boxes))
	fmt.Println("the volume of the first one is ", boxes[0].Volume(), "cm")
	fmt.Println("the color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("the biggest one is ", boxes.BiggestColor().String())
	fmt.Println("let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("the color of the second one is ", boxes[1].color.String())
	fmt.Println("boviously, now, the biggst one is ", boxes.BiggestColor().String())
}

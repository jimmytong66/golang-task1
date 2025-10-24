package main

import (
	"fmt"
	"math"
)

/* 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。 */

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64 //长
	Width  float64 //宽
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}

type Circle struct {
	Radius float64
}

func (r *Circle) Area() float64 {
	return (*r).Radius * r.Radius * math.Pi
}

func (r *Circle) Perimeter() float64 {
	return 2 * math.Pi * r.Radius
}

// 获取类型形状
func getName(s Shape) string {
	switch s.(type) {
	case *Rectangle:
		return "矩形"
	case *Circle:
		return "圆形"
	default:
		return "未知类型"
	}
}

func main() {
	// 创建不同形状

	shapes := []Shape{
		&Rectangle{
			Length: 2.0,
			Width:  4.0,
		},
		&Circle{
			Radius: 5.0,
		},
	}
	for _, s := range shapes {

		fmt.Printf("%s 面积: %.2f\n", getName(s), s.Area())
		fmt.Printf("%s 周长: %.2f\n", getName(s), s.Perimeter())

	}
	// fmt.Printf("矩形的面积: %.2f\n")

}

package main

import "fmt"

/* 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。 */

func sliceMultiTwo(slicePtr *[]int) {
	// 打印切片指针本身的值（即切片的地址）
	fmt.Printf("slicePtr 本身的地址: %p\n", &slicePtr)
	//切片指针的值
	fmt.Printf("slicePtr 存的值(即切片的地址): %p\n", slicePtr)
	fmt.Println("----------------------")

	// 打印切片的内容
	fmt.Printf("slicePtr 指向的切片内容: %v\n", *slicePtr)
	for i := range *slicePtr {

		(*slicePtr)[i] *= 2
		fmt.Println("=== 正确写法中的地址 ===")
		fmt.Printf("(*slicePtr)[%d] 地址: %p，值: %d\n", i, &(*slicePtr)[i], (*slicePtr)[i])
		// fmt.Println("=== 错误写法中的地址 ===")
		// fmt.Printf("v 地址: %p，值: %d\n", &v, v)
	}
}

func main() {
	mySlice := []int{
		1, 2, 3, 4,
	}
	//slice本质是个结构体
	fmt.Printf("mySlice[0]的地址: %p\n", &mySlice[0])
	fmt.Printf("mySlice的地址: %p\n", &mySlice)
	sliceMultiTwo(&mySlice)

	fmt.Println(mySlice)

}

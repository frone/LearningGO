package main

import (
	"fmt"
	"math/rand"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	// 使用make创建空切片
	integer := make([]int, 2)
	fmt.Println(integer)
	integer = nil
	fmt.Println(integer)

	anArray := [5]int{-1, -2, -3, -4, -5}
	// [:]操作只是将引用指向数组，并没有创建一份数组的拷贝:
	refAnArray := anArray[:]
	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -100
	fmt.Println(refAnArray)

	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println("--------------")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(split(17))
}

package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func rangeTest() {
	anArray := [5]int{0, 1, -1, 2, -2}
	// 对一个数组使用range关键字会返回两个值：一个是数组的索引，一个是该索引上的值。
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}

// 切片
// 切片作为函数的形参时是传引用操作，传递的是指向切片的内存地址，这意味着在函数中对切片的任何操作都会在函数结束后体现出来。
// 另外，函数中传递切片要比传递同样元素数量的数组高效，因为Go只是传递指向切片的内存地址，而非拷贝整个切片。
func sliceTest() {
	// aSliceLiteral := []int{1, 2, 3, 4, 5}
	integer := make([]int, 20)
	integer = append(integer, -5900)
	for _, value := range integer {
		fmt.Println(value)
	}
}

// 打印一个slice的所有元素
func printSlice(x []int) {
	for _, number := range x {
		fmt.Println(number, " ")
	}
	fmt.Println()
}

func sliceTest2() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
}

func main() {
	// d3()
	// a()
	// fmt.Println("main() ended!")
	// rangeTest()
	// sliceTest()
	sliceTest2()
}

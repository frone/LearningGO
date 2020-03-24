package main

import (
	"fmt"
	"time"
)

const s string = "constant"

func main() {
	fmt.Println("GO")
	var e int
	fmt.Println(e)
	fmt.Println(s)

	// 简单循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始化/条件/后续形式 for 循环。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 在一个 `case` 语句中，你可以使用逗号来分隔多个表
	// 达式。在这个例子中，我们很好的使用了可选的
	// `default` 分支。
	fmt.Println(time.Now())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// 不带表达式的 `switch` 是实现 if/else 逻辑的另一种
	// 方式。这里展示了 `case` 表达式是如何使用非常量的。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	// 不像数组，slice 的类型仅由它所包含的元素决定（不像
	// 数组中还需要元素的个数）。要创建一个长度非零的空
	// slice，需要使用内建的方法 `make`。这里我们创建了一
	// 个长度为3的 `string` 类型 slice（初始化为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 如你所料，`len` 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 作为基本操作的补充，slice 支持比数组更多的操作。
	// 其中一个是内建的 `append`，它返回一个包含了一个
	// 或者多个新值的 slice。注意我们接受返回由 append
	// 返回的新的 slice 值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 `copy`。这里我们创建一个空的和 `s` 有
	// 相同长度的 slice `c`，并且将 `s` 复制给 `c`。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 `slice[low:high]` 语法进行“切片”操
	// 作。例如，这里得到一个包含元素 `s[2]`, `s[3]`,
	// `s[4]` 的 slice。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 从 `s[0]` 到（但是包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中声明并初始化一个 slice 变量。
	dcl := []string{"g", "h", "i"}
	fmt.Println("dcl:", dcl)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不
	// 同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

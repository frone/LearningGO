package test

import "testing"

// 指针测试，不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T,%T", a, aPtr)
}

// String 默认初始化值是空字符串
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}

func TestCompareArray(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	b := [4]int{2, 3, 4, 5}
	c := [...]int{2, 3, 4, 5}
	// d := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b, b == c, a == c)
	// t.Log(c == d)
}

//按位清零 &^

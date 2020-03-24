package test

import "testing"

func TestArrayInit(t *testing.T) {
	a := [3]int{1, 2, 3}
	var b [3]int
	// ...表示动态获取长度
	c := [...]int{2, 3, 4, 5, 6}
	t.Log(a)
	t.Log(b)
	t.Log(c)
	// 遍历数组
	for i, value := range c {
		t.Log(i, value)
	}
	// _ 占位符
	for _, value := range c {
		t.Log(value)
	}
}

func TestSliceInit(t *testing.T) {
	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	for _, value := range s2 {
		t.Log(value)
	}

	s3 := append(s2, 6, 7, 8)
	t.Log(len(s3), cap(s3))
}

// capacity 每次增长翻倍 length +1
// slice 是共享空间，会影响通地址的切片，且不能比较
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

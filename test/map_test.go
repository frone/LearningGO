package test

import (
	"fmt"
	"math"
	"testing"
)

func TestMap(t *testing.T) {
	// map [key_type]value_type{}
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log(m1["a"])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// 空的map值默认为0
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	// 使用if变量赋值
	if v, ok := m1[3]; ok {
		t.Log("Value of Key 3 is existing", v)
	} else {
		t.Log("Value of Key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

	m2 := map[int]func(op, exp float64) float64{}
	m2[1] = func(op, exp float64) float64 { return math.Pow(op, exp) }
	t.Log(m2[1](2, 3))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	CheckKeyInMap(mySet, 1)
	mySet[3] = true
	t.Log(len(mySet))
	// 删除map的一个key
	delete(mySet, 1)
	CheckKeyInMap(mySet, 1)
}

func CheckKeyInMap(mySet map[int]bool, n int) {
	if mySet[n] {
		fmt.Printf("%d is existing!", n)
	} else {
		fmt.Printf("%d is not existing!", n)
	}
}

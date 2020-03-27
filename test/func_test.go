package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 在目标内部函数外边包装一层计时函数，并返回
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可以参数函数
func Sum(ops ...int) int {
	ret := 0
	for _, value := range ops {
		ret += value
	}
	return ret
}

func TestVariableParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(11, 12, 3, 4, 5))
}

// 延迟执行 defer
func TestDefer(t *testing.T) {
	// 使用defer延迟运行一个匿名函数
	defer func() {
		t.Log("Cleaning resource")
	}()
	t.Log("Started")
	// 程序异常中断
	panic("Fatal Error!")
	// defer仍然运行
}

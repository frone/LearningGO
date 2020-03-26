package test

import (
	"runtime"
	"testing"
)

// 连续常量 简洁赋值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestFirstTry(t *testing.T) {
	t.Log("My First try!")
}

func TestConstant(t *testing.T) {
	t.Log(Monday, Wednesday)
}

func TestConstant2(t *testing.T) {
	t.Log(Readable, Executable)
}

// 不需要break
func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "linux":
		t.Log("linux")
	default:
		t.Logf("%s", os)
	}
}

func TestNumber(t *testing.T) {
	for i := 1; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		}
	}
}

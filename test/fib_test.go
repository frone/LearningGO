package test

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	t.Log(" ", a)
	for i := 0; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}
func TestExange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}


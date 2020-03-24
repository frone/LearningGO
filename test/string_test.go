package test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, value := range parts {
		t.Log(value)
	}
	j := strings.Join(parts, "-")
	t.Log(j)
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)
	//转换为数字时会附带错误值
	if i, err := strconv.Atoi("10a"); err == nil {
		t.Log(5 + i)
	} else {
		t.Log(err)
	}

}

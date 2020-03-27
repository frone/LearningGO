package interface_test

import (
	"fmt"
	"testing"
)

type Progammer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Progammer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}

// 空接口测试
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Printf("unknown - %T", p)
}

func DoSomethingV2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Printf("unknown - %T", v)
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomethingV2(100)
	DoSomethingV2("abc")
	DoSomethingV2([4]int{})
}

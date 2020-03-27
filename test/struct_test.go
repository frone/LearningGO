package test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// func (e Employee) string() string {
// 	return fmt.Sprintf("ID:%s - NAME:%s - AGE:%d", e.Id, e.Name, e.Age)
// }

// 使用指针没有内容复制，不占用新的内存空间
func (e *Employee) string() string {
	return fmt.Sprintf("ID:%s - NAME:%s - AGE:%d", e.Id, e.Name, e.Age)
}
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"001", "Bob", 55}
	e1 := Employee{Name: "Jack", Age: 30}
	e2 := new(Employee)
	e2.Id = "003"
	e2.Age = 26
	e2.Name = "David"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e1.Id == "")
	t.Logf("e is %T", e)
	t.Logf("&e is %T", &e)
	t.Logf("e2 is %T", e2)
	t.Log(e.string())
}

package interface_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// DOG类型有一个内部成员变量 PET
type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("WWWW")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("WANG", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("CHAO")
}

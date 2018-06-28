package main

import (
	"fmt"
)

type Far struct {
	name string
}
type Fay struct {
	Far
}

type Per interface {
	act()
}

//封装继承多肽
func main() {
	// p := Fay{Far: Far{name: "parent"}}
	// fmt.Println(p.name)
	// p.fly()
	var p Fay
	var d Far
	p.act()
	d.act()
}

func (par Far) fly() {
	fmt.Println("封装")
}

func (par Far) act() {
	fmt.Println("duotai1")
}
func (par Fay) act() {
	fmt.Println("duotai2")
}

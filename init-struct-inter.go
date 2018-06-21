package main

import (
	"fmt"
)

type Per struct {
	a string
	b string
}

type Inter interface {
	fly()
}

func main() {
	c := new(Per)
	fmt.Println(c)

	var d Inter = new(Per)
	d.fly()
	//fmt.Println(e)
}

func fly() {
	fmt.Println("flyying")
}

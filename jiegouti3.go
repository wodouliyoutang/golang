package main

import (
	"fmt"
)

type Fly struct {
	Name string
	age  int
}

func main() {
	a := Fly{"dabao", 8}
	a.Act()

}

func (fly *Fly) Act() {
	fmt.Println("success")
}

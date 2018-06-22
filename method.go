package main

import "fmt"

type A struct {
	name int
}

func main() {
	a := A{}
	a.print(3)
	fmt.Println(a.name)
}

func (a *A) print(b int) {
	a.name = 2 + b
	fmt.Println(*a)
}

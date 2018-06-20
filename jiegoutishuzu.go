package main

import "fmt"

//结构体与数组结合
type Sudo struct {
	a int
	b string
}

type Jor struct {
	name string
	sex  int
}

func main() {
	//将结构体嵌入数组中
	c := []Sudo{
		{3, "y"},
		{5, "i"},
	}
	fmt.Println(c)

	//将数组放入结构体中
	d := struct {
		name []Sudo
	}{c}
	fmt.Println(d)

	// example
	arr := []Jor{
		{"tom", 1},
		{"lily", 0},
	}
	fmt.Println(arr)

	stry := struct {
		test []Jor
	}{arr}
	fmt.Println(stry)
}

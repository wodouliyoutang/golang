package main 

import "fmt"

type PEOPLE interface{
	name() string
	connect()
}

type men struct{
	age int
	sex string
}

func main() {
	var c PEOPLE
	c = men{18,"男"}
	c.connect()
}

func (a men) name() string {
	a.sex = "男"
	return a.sex
}

func (b men) connect(){
	fmt.Println(b)
}
package main 

import "fmt"

type humen struct{
	sex int
}

type men struct{
	humen		//继承结构体
	age int
	name struct{
		ll int
	}
}

func main() {
	a := men{age:3,humen:humen{sex:9}}
	a.name.ll =32
	fmt.Println(a)	

}
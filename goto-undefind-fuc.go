package main

import (
	"fmt"
)

func main() {
	/*
			br := 3
		HERE:
			fmt.Println(br)
			fmt.Println("success")
			br++
			if br < 5 {
				goto HERE
			}
	*/

	var arr = []int{1, 3, 9}
	str := "success"
	undef(str, arr)

}

//不定传参,且不确定类型
func undef(str string, arg ...interface{}) {
	for _, v := range arg {
		fmt.Println(v)
	}

	fmt.Println(str)

}

package main

import (
	"fmt"
	_ "reflect"
)

func main() {
	/*
		// 双目运算符
		var mast = 8 << 2  // 左移 2   => 8*(2*2)	=>32
		var masts = 8 >> 3 //右移 2	  => 8 /(2*3)	=>1
		fmt.Println(mast)
		fmt.Println(masts)
	*/

	// b := "哈喽world"
	// for k, v := range b {
	// 	fmt.Println(k, v)
	// }
	//
	// string()中放 ascii 码可转化相应字符串
	s := "哈喽world"
	fmt.Println([]rune(s))
	fmt.Println([]rune(s)[3])
	s = string([]rune(s)[3])
	fmt.Println(s)
	fmt.Println(string(21949))

}

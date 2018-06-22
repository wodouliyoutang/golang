package main

import (
	"fmt"
)

// 不管任何下 都执行的一个东西
// 遵循后执行
func main() {
	/*
	  fmt.Println("1")
	  defer fmt.Println("2")
	  defer fmt.Println("3")

	  for i := 0; i<3; i++ {
	      defer func(){
	        fmt.Println(i)
	      }()
	    fmt.Println(i)
	  }
	*/
	var arr = []int{1, 4, 7}
	// a := func() bool {
	// 	for _, v := range arr {
	// 		if v == 4 {
	// 			return true
	// 		}
	// 		return false
	// 	}
	// }()
	// b := a
	// fmt.Println(b)

}

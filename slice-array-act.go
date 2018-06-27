package main

import "fmt"

func main() {
	mySlice2 := []int{8, 9, 10}
	mySlice1 := []int{5, 6, 7}
	mySlice3 := append(mySlice1, mySlice2...)
	fmt.Println(mySlice3)

}

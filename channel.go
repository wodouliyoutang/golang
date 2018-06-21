package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 6, 9, 7}
	resultChan := make(chan int, 2)
	go sun(arr[:len(arr)/2], resultChan)
	go sun(arr[len(arr)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println(sum1, sum2)
}

func sun(arr []int, resultChan chan int) {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	resultChan <- sum
}

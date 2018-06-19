package main

import (
	"fmt"
	_ "time"
)

//并发测试  主线程中赋值,就会阻塞主线程,知道值被取出
func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		fmt.Println("son")
	}()
	<-ch
	fmt.Println("par")
}

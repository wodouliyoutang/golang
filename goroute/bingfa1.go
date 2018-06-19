package main

import (
	"fmt"
)

//有缓冲的chan
func main() {
	ch := make(chan string, 2)
	ch <- "channel1"
	ch <- "channel2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

package main

import (
	"fmt"
)

func main() {
	var line = []string{"a", "b", "c"}
	i := 0
	for {
		fmt.Println(line[i])
		i++
		if i > 2 {

			break
		}

	}

}

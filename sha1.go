package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

type Wxpar struct {
	Signature string
	Timestamp string
	Nonce     string
	Echostr   string
}

func main() {
	str := "aabbcc1223"
	t := sha1.New()
	io.WriteString(t, str)
	a := fmt.Sprintf("%x", t.Sum(nil))
	fmt.Println(a)
}

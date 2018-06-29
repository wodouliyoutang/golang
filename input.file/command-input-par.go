package main

import "flag"
import "fmt"

//声明一个string flag  参数  1.参数输入名称 2.默认值  3.说明
var a *string = flag.String("a", "amorenzhi", "dakaiafile")
var b *string = flag.String("b", "bmorenzhi", "dakaibfile")
var c *string = flag.String("c", "cmorenzhi", "dakaicfile")

func main() {
	flag.Parse() //解析命令行输入的值
	if a != nil {
		if b != nil {
			if c != nil {
				fmt.Println("a =", *a, "b =", *b, "c =", *c)
			}
		}
	}
}

//readme
//  ①编译 go build command-input-par.go
//  ②执行 command-input-par.exe -a a.ini -b b.ini -c c.ini

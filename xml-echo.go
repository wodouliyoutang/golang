package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

//结构体后的自定义标签 需要在使用指定的包才生效，比如 xml json 包
type xmls struct {
	XMLName xml.Name `xml:"xmltitle"` //xml 的根节点
	AA      string   //`xml:",attr"`
	BB      string   //`xml:"_"`
	CC      string   //`xml:"name,attr"`
	DD      string   //`xml:",innerxml"`
	EE      string   //`xml:",chardata"`
}

func main() {
	v := &xmls{}
	v.AA = "valA"
	v.BB = "valB"
	v.CC = "valC"
	v.DD = "valD"
	v.EE = "valE"
	output, err := xml.MarshalIndent(&v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output)) //两种都可以输出
	os.Stdout.Write(output)
}

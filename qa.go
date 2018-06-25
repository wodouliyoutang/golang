package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}

func main() {
	content, err := ioutil.ReadFile("file/studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(result.ResourceString)
}

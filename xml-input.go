package main

import (
	"encoding/xml"
	"fmt"
)

type Gypic struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int
}

func main() {
	wexml := `<xml><ToUserName>toUser</ToUserName><FromUserName>formUser</FromUserName><CreateTime>1348831860</CreateTime><MsgType>text</MsgType><Content>hello world</Content><MsgId>1234567890123456</MsgId></xml>`
	result := Gypic{}
	xml.Unmarshal([]byte(wexml), &result)
	fmt.Println(result.CreateTime)

}

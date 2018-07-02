package main

import (
	"fmt"
	"net/http"
	"user"
	"wx"
)

/**
 * 主函数路由定义
 * @return {[type]} [description]
 */
func main() {
	http.HandleFunc("/", func(Response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Response, "首页")
	})

	http.HandleFunc("/login", user.Login)

	http.HandleFunc("/index", user.Index)
	http.HandleFunc("/string-log", user.StringLog)
	http.HandleFunc("/upload", user.Upload)
	http.HandleFunc("/json", user.JsonEco)
	http.HandleFunc("/xml", user.XmlAct)
	http.HandleFunc("/wxsign", wx.Wxsign)
	http.HandleFunc("/database", sqlact.Cre)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("Println")         //服务端
	// fmt.Fprint(rw, "hello worlds") //客户端

}

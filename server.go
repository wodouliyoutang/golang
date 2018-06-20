package main

import (
	"fmt"
	"net/http"
	"user"
)

/**
 * 主函数路由定义
 * @return {[type]} [description]
 */
func main() {
	http.HandleFunc("/", func(Response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Response, "/index")
	})

	http.HandleFunc("/login", user.Login)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("Println")         //服务端
	// fmt.Fprint(rw, "hello worlds") //客服端

}

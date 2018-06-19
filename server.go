package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template

func main() {
	http.HandleFunc("/", func(Response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			initTemplate("C:\\Users\\datu\\go\\src\\user\\login.html")
			//template.ParseFiles("C:\\Users\\datu\\go\\src\\user\\login.html")
			fmt.Println("server")
		}
	})
	http.HandleFunc("/index", func(Response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(Response, "/index")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 	if r.Method == "POST" {
	// 		var (
	// 			key   string = r.PostFormValue("key")
	// 			value string = r.PostFormValue("value")
	// 		)
	// 		fmt.Printf("key is  : %s\n", key)
	// 		fmt.Printf("value is: %s\n", value)
	// 	}
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Println")         //服务端
	// fmt.Fprint(rw, "hello worlds") //客服端

}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

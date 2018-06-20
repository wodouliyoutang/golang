package user

import (
	"fmt"
	"html/template"
	"net/http"
	_ "reflect"
	"strings"
)

type Data struct {
	Errorinfo string
}

// 一个大坑 大小写结构体变量-------------
func Login(response http.ResponseWriter, request *http.Request) {
	var data Data
	if request.Method == "POST" {
		user := request.PostFormValue("name")
		pwd := request.PostFormValue("pwd")
		idcard := request.PostFormValue("idcard")
		if strings.Count(user, "") < 20 && strings.Count(user, "") > 7 {
			if strings.Count(pwd, "") > 7 && strings.Count(pwd, "") < 21 {
				if strings.Count(idcard, "") > 17 && strings.Count(idcard, "") < 20 {
					fmt.Println("hello worlds")
					data.Errorinfo = "hello worlds"
				}
				data.Errorinfo = "idcard-error"
				fmt.Println("card")
			}
			data.Errorinfo = "pwd-error"
			fmt.Println("pwd")
		}
		data.Errorinfo = "user-error"
		fmt.Println("user")

	}

	tplhtml := template.Must(template.ParseFiles("C:\\Users\\datu\\go\\src\\user\\login.html"))
	tplhtml.Execute(response, data)
}

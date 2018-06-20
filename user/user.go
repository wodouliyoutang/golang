package user

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	_ "reflect"
	"strings"
)

type Data struct {
	Errorinfo string
}

/**
* 渲染模板,传递的数据的值命名要大写,要以结构体传递
 */
func Login(response http.ResponseWriter, request *http.Request) {
	var data Data
	if request.Method == "POST" {
		user := request.PostFormValue("name")
		pwd := request.PostFormValue("pwd")
		idcard := request.PostFormValue("idcard")
		if strings.Count(user, "") < 20 && strings.Count(user, "") > 7 {
			if strings.Count(pwd, "") > 7 && strings.Count(pwd, "") < 21 {
				if strings.Count(idcard, "") > 17 && strings.Count(idcard, "") < 20 {
					fmt.Println("successr")
					http.Redirect(response, request, "/index", http.StatusFound)
				} else {
					data.Errorinfo = "idcard-error"
					fmt.Println("card")
				}
			} else {
				data.Errorinfo = "pwd-error"
				fmt.Println("pwd")
			}
		} else {
			data.Errorinfo = "user-error"
			fmt.Println("user")
		}

	}

	tplhtml := template.Must(template.ParseFiles("user/login.html"))
	tplhtml.Execute(response, data)
}

func Index(response http.ResponseWriter, request *http.Request) {
	data := struct {
		Info string
	}{
		Info: "weclome",
	}

	tplhtml := template.Must(template.ParseFiles("user/index.html"))
	tplhtml.Execute(response, data)
}

func Dologin(response http.ResponseWriter, request *http.Request) {
	//   io.Writestring
	//   log.Fatal
	//   字符串 + 号
	//   request.URL.string()
	//   声明方法如何调用?接受体声明了,且都符合,正常调用了
	//   src 下多项目多文件使用
	fmt.Println(request.PostFormValue("test"))
}

func upload(response http.ResponseWriter, request *http.Request) {
	message := ""
	if request.Method == "POST" {
		request.ParseMultipartForm(32 << 20)
		file, hand, err := request.FormFile("filename")
		if err != nill {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nill {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		message := "upload-success"
	}
}

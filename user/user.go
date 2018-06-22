package user

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	_ "log"
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
/**
 * 登录操作
 * @param {[type]} response http.ResponseWriter [description]
 * @param {[type]} request  *http.Request       [description]
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

/**
 * 视图层数据推送
 * 传递的数据的值命名要大写
 * @param {[type]} response http.ResponseWriter [description]
 * @param {[type]} request  *http.Request       [description]
 */
func Index(response http.ResponseWriter, request *http.Request) {
	data := struct {
		Info string
	}{
		Info: "weclome",
	}

	tplhtml := template.Must(template.ParseFiles("user/index.html"))
	tplhtml.Execute(response, data)
}

func StringLog(response http.ResponseWriter, request *http.Request) {

	a := "字符串拼接"
	io.WriteString(response, "输出到客户端"+a)         //输出客户端
	io.WriteString(response, request.RequestURI) //当前url

	//日志形式打印到终端
	log.Fatal("success")

	log.Print("success")
	os.Exit(1) //退出当前执行程序

}

/**
 * 上传文件并判断是否存在
 * @param {[type]} response http.ResponseWriter [description]
 * @param {[type]} request  *http.Request       [description]
 */
func Upload(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		request.ParseMultipartForm(32 << 20)
		point, fileinfo, err := request.FormFile("fil")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer point.Close()

		filepath := "C:\\Users\\datu\\go\\src\\file"
		infos, err := ioutil.ReadDir(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		isfile := func() bool {
			for _, v := range infos {
				if v.Name() == fileinfo.Filename {
					return true
				}
			}
			return false
		}()
		if isfile {
			http.Error(response, "该文件已存在", http.StatusInternalServerError)
			return
		}
		//f, err := os.OpenFile("./file/"+fileinfo.Filename, os.O_WRONLY|os.O_CREATE, 0666) //两种创建文件方法
		f, err := os.Create("./file/" + fileinfo.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, point)
	}
}

package wx

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	_ "reflect"
	"sort"
)

/**
开发者提交信息后，微信服务器将发送GET请求到填写的服务器地址URL上，GET请求携带参数如下表所示：
signature	微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
timestamp	时间戳
nonce	随机数
echostr	随机字符串
将token、timestamp、nonce三个参数进行字典序排序 2）将三个参数字符串拼接成一个字符串进行sha1加密 3）开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
*/
const (
	APPID = "wx9f04596fca7b9105"
	Token = "datuhongan_travel"
)

type Wxpar struct {
	Signature string
	Timestamp string
	Nonce     string
	Echostr   string
}

func Wxsign(Response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("echostr") == "" {
		fmt.Fprintln(Response, "response")
	} else {
		wxpar := new(Wxpar)
		wxpar.Signature = request.URL.Query().Get("signature")
		wxpar.Timestamp = request.URL.Query().Get("timestamp")
		wxpar.Nonce = request.URL.Query().Get("nonce")
		wxpar.Echostr = request.URL.Query().Get("echostr")

		sort_data := sort.StringSlice{Token, wxpar.Timestamp, wxpar.Nonce}

		sort.Strings(sort_data)
		str := ""
		for _, v := range sort_data {
			str += v
		}
		s := sha1.New()
		io.WriteString(s, str)
		a := fmt.Sprintf("%x", s.Sum(nil))
		// fmt.Fprint(Response, wxpar.Echostr)
		fmt.Println(a)
		fmt.Println(wxpar.Echostr)
		// if a == wxpar.Signature {
		// 	fmt.Fprint(Response, a)
		// }

	}

}

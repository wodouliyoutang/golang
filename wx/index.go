package wx

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "reflect"
	"sort"
	"time"
)

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

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

func Wxsign(Response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("echostr") == "" {
		request.ParseForm()
		if request.Method == "POST" {
			textRequestBody := parseTextRequestBody(request)
			if textRequestBody != nil {
				// fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				// 	textRequestBody.Content,
				// 	textRequestBody.FromUserName)

				responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
					textRequestBody.FromUserName,
					textRequestBody.Content)
				if err != nil {
					log.Println("Wechat Service: makeTextResponseBody error: ", err)
					return
				}
				fmt.Fprintf(Response, string(responseTextBody))
			}
		}

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
		if a == wxpar.Signature {
			fmt.Fprint(Response, wxpar.Echostr)
		}

	}

}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("接受体:", string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = fromUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = "请输入helloworld"
	if content == "helloworld" {
		textResponseBody.Content = "golang-web"
	}
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

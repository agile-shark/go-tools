package wechat

import (
	"fmt"
	"log"
	"net/http"
)

//func makeSignature(timestamp, nonce string) string {
//	sl := []string{token, timestamp, nonce}
//	sort.Strings(sl)
//	s := sha1.New()
//	io.WriteString(s, strings.Join(sl, ""))
//	return fmt.Sprintf("%x", s.Sum(nil))
//}

//func validateUrl(w http.ResponseWriter, r *http.Request) bool {
//	timestamp := strings.Join(r.Form["timestamp"], "")
//	nonce := strings.Join(r.Form["nonce"], "")
//	signatureGen := makeSignature(timestamp, nonce)
//
//	signatureIn := strings.Join(r.Form["signature"], "")
//	if signatureGen != signatureIn {
//		return false
//	}
//	echostr := strings.Join(r.Form["echostr"], "")
//	fmt.Fprintf(w, echostr)
//	return true
//}

//func parseTextRequestBody(r *http.Request) *TextRequestBody {
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		log.Fatal(err)
//		return nil
//	}
//	fmt.Println(string(body))
//	requestBody := &TextRequestBody{}
//	xml.Unmarshal(body, requestBody)
//	return requestBody
//}
//
//func value2CDATA(v string) CDATAText {
//	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
//	return CDATAText{"<![CDATA[" + v + "]]>"}
//}
//
//func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
//	textResponseBody := &TextResponseBody{}
//	textResponseBody.FromUserName = value2CDATA(fromUserName)
//	textResponseBody.ToUserName = value2CDATA(toUserName)
//	textResponseBody.MsgType = value2CDATA("text")
//	textResponseBody.Content = value2CDATA(content)
//	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
//	return xml.MarshalIndent(textResponseBody, " ", "  ")
//}

func processRequestMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}

	if r.Method == "POST" {
		textRequestBody := parseTextRequestBody(r)
		if textRequestBody != nil {
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)
			responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				"Hello, "+textRequestBody.FromUserName)
			if err != nil {
				log.Println("Wechat Service: makeTextResponseBody error: ", err)
				return
			}
			w.Header().Set("Content-Type", "text/xml")
			fmt.Println(string(responseTextBody))
			fmt.Fprintf(w, string(responseTextBody))
		}
	}
}

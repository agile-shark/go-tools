package main

import (
	"github.com/agile-shark/go-tools/wechat"
	"log"
	"fmt"
	"strings"
	"strconv"
	"github.com/agile-shark/go-tools/email"
)

func email_alert(content string)  (error error){
	emails := email.NewEmail("2424751761@qq.com;", "人民在线", content)
	error = email.SendEmail(emails, nil)
	return error
}

func wechat_alert(content string)  error{

	// Fetch access_token
	accessToken, expiresIn, error := wechat.FetchAccessToken()
	if error != nil {
		log.Println("Get access_token error:", error)
		return error
	}
	fmt.Println(accessToken, expiresIn)

	//Post custom service message
	openID := "oT6jJvqQlsQDUmnN_rMzHaV5pD-8"
	error = wechat.PushCustomMsg(accessToken, openID, content)
	if error != nil {
		log.Println("Push custom service message err:", error)
		return error
	}
	return nil
}

func main()  {
	message_count := 10000
	alert_message := strings.Join([]string{"消息队列", "83", "消息量超过预警阀值 \n 当前消息总量为", strconv.Itoa(message_count)}, "")
	email_alert(alert_message)
	wechat_alert(alert_message)
}
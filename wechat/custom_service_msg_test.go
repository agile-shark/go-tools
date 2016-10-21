package wechat

import (
	"testing"
	"log"
	"fmt"
)

func Test_send(t *testing.T) {
	// Fetch access_token
	accessToken, expiresIn, err := FetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)

	//Post custom service message
//	openID := "oT6jJvqQlsQDUmnN_rMzHaV5pD-8"
//	msg := "你好" + "\U0001f604"
//	err = PushCustomMsg(accessToken, openID, msg)
//	if err != nil {
//		log.Println("Push custom service message err:", err)
//		return
//	}
}

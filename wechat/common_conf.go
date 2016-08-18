package wechat

import (
	"encoding/xml"
	"time"
)

const (
	appID                = "wxaeb7504679f0d461"
	appSecret            = "a56163c3c7825c543a7262976d9009fc"
	accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/mass/send"

	token          = "wechat4go"
	encodingAESKey = "kZvGYbDKbtPbhv4LBWOcdsp5VktA3xe9epVhINevtGg"
)


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
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}


type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

type CustomServiceMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	Text    TextMsgContent `json:"text"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}


//type TextRequestBody struct {
//	XMLName      xml.Name `xml:"xml"`
//	ToUserName   string
//	FromUserName string
//	CreateTime   time.Duration
//	MsgType      string
//	Content      string
//	MsgId        int
//}

type TextResponseBody1 struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   string
	MsgType      CDATAText
	Content      CDATAText
}

type EncryptRequestBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

type EncryptResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATAText
	MsgSignature CDATAText
	TimeStamp    string
	Nonce        CDATAText
}

type EncryptResponseBody1 struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
}

//type CDATAText struct {
//	Text string `xml:",innerxml"`
//}
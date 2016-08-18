package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchAccessToken() (string, float64, error) {
	requestLine := strings.Join([]string{accessTokenFetchUrl, "?grant_type=client_credential&appid=", appID, "&secret=", appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", 0.0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0.0, err
	}

	fmt.Println(string(body))

	//Json Decoding
	if bytes.Contains(body, []byte("access_token")) {
		fmt.Println("return ok")
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			return "", 0.0, err
		}
		return atr.AccessToken, atr.ExpiresIn, nil
	} else {
		fmt.Println("return err")
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			return "", 0.0, err
		}
		return "", 0.0, fmt.Errorf("%s", ater.Errmsg)
	}
}

func PushCustomMsg(accessToken, toUser, msg string) error {
	csMsg := &CustomServiceMsg{
		ToUser:  toUser,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(csMsg, " ", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	postReq, err := http.NewRequest("POST", strings.Join([]string{customServicePostUrl, "?access_token=", accessToken}, ""), bytes.NewReader(body))
	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

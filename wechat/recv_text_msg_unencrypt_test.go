package wechat

import (
	"testing"
	"net/http"
	"log"
)

func Test_ListenAndServe(t *testing.T) {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", processRequestMsg)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}

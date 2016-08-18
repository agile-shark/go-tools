package wechat

import (
	"testing"
	"log"
	"net/http"
)

func Test_ListenAndServe_1(t *testing.T) {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}

package email

import(
    "testing"
    "fmt"
)

func Test_send_email(t *testing.T) {
    content := " my dear令"
    email := NewEmail("xxxxxxx@qq.com;", "test golang email", content)
    sMTPServer := &SMTPServer{ServerAddr : SERVER_ADDR, Host : HOST, User : "*********@gmail.com", Password : "*********"}
    err := SendEmail(email, sMTPServer)
    fmt.Println(err)
}
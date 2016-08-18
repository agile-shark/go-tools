package email

import (
    "net/smtp"
    "strings"
)

const (
    HOST        = "smtp.163.com"
    SERVER_ADDR = "smtp.163.com:25"
    USER        = "*****@163.com"         //发送邮件的邮箱
    PASSWORD    = "*************"         //发送邮件邮箱的密码
)

type SMTPServer struct {
    Host        string
    ServerAddr  string
    User        string  //发送邮件的邮箱
    Password    string  //发送邮件邮箱的密码
}

type Email struct {
    to      string "to"
    subject string "subject"
    msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
    return &Email{to: to, subject: subject, msg: msg}
}

func SendEmail(email *Email, server *SMTPServer) error {

    auth := smtp.PlainAuth("", server.User, server.Password, server.Host)
    sendTo := strings.Split(email.to, ";")
    done := make(chan error, 1024)

    go func() {
        defer close(done)
        for _, v := range sendTo {
            str := strings.Replace("From: "+server.User+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg
            err := smtp.SendMail(
                server.ServerAddr,
                auth,
                server.User,
                []string{v},
                []byte(str),
            )
            done <- err
        }
    }()
    for i := 0; i < len(sendTo); i++ {
        <-done
    }
    return nil
}
package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func Send(fileName string) {
	auth = smtp.PlainAuth("", "email", "pass", "smtp.gmail.com")
	templateData := struct {
		FileName string
	}{
		FileName: fileName,
	}
	r := newRequest([]string{"emailparaenvio"}, "Hello Daniel!", "Hello, World!")
	err := r.ParseTemplate("./email/template.html", templateData)
	if err != nil {
		fmt.Println(err)
		return
	}
	ok, err := r.sendEmail()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(ok)
}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func newRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) sendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "email", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

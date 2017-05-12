package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var auth smtp.Auth

func Send(fileName string) {

	from := mail.NewEmail("Colossus", "daniel.viana@m4u.com.br")
	subject := "Novos Arquivos da TIM"
	to := mail.NewEmail("Daniel", "daniel.viana@m4u.com.br")
	templateData := struct {
		FileName string
	}{
		FileName: fileName,
	}
	err, buf := ParseTemplate("./email/template.html", templateData)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := mail.NewContent("text/html", buf.String())
	m := mail.NewV3MailInit(from, subject, to, content)
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func newRequest(to []string, from string, subject string) *Request {
	return &Request{
		from:    from,
		to:      to,
		subject: subject,
	}
}
func ParseTemplate(templateFileName string, data interface{}) (error, *bytes.Buffer) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err, nil
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err, nil
	}
	return nil, buf
}

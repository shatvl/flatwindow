package services

import (
	"bytes"
	"github.com/kataras/go-mailer"
	"html/template"
)

type Mailer interface {
	SendMail() error
}

type Smtp struct {
	Mailer *mailer.Mailer
}

func parseTemplate(fileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(fileName)

	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)

	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func NewSmtpMailer() *Smtp {
	// sender configuration.
	config := mailer.Config{
		Host:     "smtp.gmail.com",
		Username: "shatilenya95@gmail.com",
		FromAddr: "shatilenya95@gmail.com",
		Port:     587,
		// Enable UseCommand to support sendmail unix command,
		// if this field is true then Host, Username, Password and Port are not required,
		// because these info already exists in your local sendmail configuration.
		//
		// Defaults to false.
		UseCommand: false,
	}

	sender := mailer.New(config)

	return &Smtp{Mailer: sender}
}

func (smtp *Smtp) SendBidRequestMail() error {
	// the subject/title of the e-mail.
	subject := "Заявка на детальный запрос / просмотр"
	// the rich message body.
	// content := `<h1>Hello</h1> <br/><br/> <span style="color:red"> This is the rich message body </span>`

	// the recipient(s).
	to := []string{"vladislav.shatilenya@gmail.com", "maxxporoshin@gmail.com"}
	body, err := parseTemplate("templates/mails/bid.html", map[string]string{"username": "Vlad"})

	if err != nil {
		return err
	}

	// send the e-mail.
	err = smtp.Mailer.Send(subject, body, to...)

	if err != nil {
		println("error while sending the e-mail: " + err.Error())
	}

	return err
}

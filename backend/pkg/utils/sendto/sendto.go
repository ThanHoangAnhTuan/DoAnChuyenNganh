package sendto

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
)

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func BuildMessage(email Email) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", email.From)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", email.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", email.Body)

	return msg
}

func SendEmailOTP(to []string, templateName string, templateData map[string]interface{}) error {
	htmlBody, err := getEmailTemplate(templateName, templateData)
	if err != nil {
		return err
	}

	return send(to, htmlBody)
}

func getEmailTemplate(templateName string, templateData map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(templateName).ParseFiles("templates_custom/email/" + templateName))
	err := t.Execute(htmlTemplate, templateData)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, htmlTemplate string) error {
	globalEmail := global.Config.Email
	contentEmail := Email{
		From:    globalEmail.User,
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageEmail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", globalEmail.User, globalEmail.Password, globalEmail.Host)

	err := smtp.SendMail(globalEmail.Host+":"+globalEmail.Port, auth, globalEmail.User, to, []byte(messageEmail))
	if err != nil {
		return fmt.Errorf("send OTP to email failed: %s", err)
	}
	global.Logger.Info("Send OTP to email success")
	return nil
}

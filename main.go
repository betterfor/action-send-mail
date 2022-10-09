package main

import (
	"crypto/tls"
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("INPUT_")
}

func main() {
	//viper.ReadInConfig()

	fmt.Println(os.Environ())

	fmt.Println(viper.AllKeys())
	mailHost := viper.GetString("server_address")
	mailPort := viper.GetInt("server_port")
	secure := viper.GetBool("secure")
	username := viper.GetString("username")
	password := viper.GetString("password")
	subject := viper.GetString("subject")
	to := strings.Split(viper.GetString("to"), ",")
	from := viper.GetString("from")
	body := viper.GetString("body")
	htmlBody := viper.GetString("html_body")
	cc := strings.Split(viper.GetString("cc"), ",")
	bcc := strings.Split(viper.GetString("bcc"), ",")
	replyTo := viper.GetString("reply_to")
	inReplyTo := viper.GetString("in_reply_to")
	//ignoreCert := viper.GetBool("ignore_cert")
	convertMarkdown := viper.GetBool("convert_markdown")
	attachments := strings.Split(viper.GetString("attachments"), ",")
	priority := viper.GetString("priority")

	m := gomail.NewMessage()
	m.SetHeader("Subject", subject)
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Cc", cc...)
	m.SetHeader("Bcc", bcc...)
	m.SetHeader("Reply-To", replyTo)
	m.SetHeader("In-Reply-To", inReplyTo)
	m.SetHeader("X-Priority", priority)

	for _, a := range attachments {
		m.Attach(a)
	}

	if !secure {
		secure = mailPort == 465
	}

	if body != "" {
		m.SetBody("text/plain", getBody(body, false))
	}
	if htmlBody != "" {
		m.SetBody("text/html", getBody(htmlBody, convertMarkdown))
	}

	dialer := gomail.NewDialer(mailHost, mailPort, username, password)
	if secure {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := dialer.DialAndSend(m)
	if err != nil {
		fmt.Println(time.Now(), "send mail error: ", err)
	}
}

func getBody(bodyOrFile string, convertMarkdown bool) string {
	var body = bodyOrFile
	if strings.HasPrefix(bodyOrFile, "file://") {
		file := strings.TrimPrefix(bodyOrFile, "file://")
		bodyBts, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("attach file error: ", err)
			return err.Error()
		}
		body = string(bodyBts)
	}

	if convertMarkdown {
		body = string(blackfriday.MarkdownBasic([]byte(body)))
	}

	return body
}

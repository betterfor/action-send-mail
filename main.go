package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/russross/blackfriday"
	"gopkg.in/gomail.v2"
)

func main() {
	mailHost := os.Getenv("INPUT_SERVER_ADDRESS")
	mailPort := convertInt(os.Getenv("INPUT_SERVER_PORT"))
	secure := convertBool(os.Getenv("INPUT_SECURE"))
	username := os.Getenv("INPUT_USERNAME")
	password := os.Getenv("INPUT_PASSWORD")
	subject := os.Getenv("INPUT_SUBJECT")
	to := strings.Split(os.Getenv("INPUT_TO"), ",")
	from := os.Getenv("INPUT_FROM")
	body := os.Getenv("INPUT_BODY")
	htmlBody := os.Getenv("INPUT_HTML_BODY")
	cc := strings.Split(os.Getenv("INPUT_CC"), ",")
	bcc := strings.Split(os.Getenv("INPUT_BCC"), ",")
	replyTo := os.Getenv("INPUT_REPLY_TO")
	inReplyTo := os.Getenv("INPUT_IN_REPLY_TO")
	//ignoreCert := viper.GetBool("ignore_cert")
	convertMarkdown := convertBool(os.Getenv("INPUT_CONVERT_MARKDOWN"))
	attachments := strings.Split(os.Getenv("INPUT_ATTACHMENTS"), ",")
	priority := os.Getenv("INPUT_PRIORITY")

	m := gomail.NewMessage()
	m.SetHeader("Subject", subject)
	m.SetHeader("From", getFrom(from, username))
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

func convertBool(b string) bool {
	b = strings.ToLower(b)
	return b == "true" || b == "yes" || b == "1"
}

func convertInt(i string) int {
	ints, _ := strconv.Atoi(i)
	return ints
}

func getFrom(from, username string) string {
	ok, _ := regexp.MatchString(`.+ <.+@.+>`, from)
	if ok {
		return from
	}
	return fmt.Sprintf("%s <%s>", from, username)
}

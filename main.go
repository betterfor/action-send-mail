package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
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
	username := os.Getenv("INPUT_USERNAME")
	password := os.Getenv("INPUT_PASSWORD")
	subject := os.Getenv("INPUT_SUBJECT")
	to := os.Getenv("INPUT_TO")
	from := os.Getenv("INPUT_FROM")
	body := os.Getenv("INPUT_BODY")
	htmlBody := os.Getenv("INPUT_HTML_BODY")
	cc := os.Getenv("INPUT_CC")
	bcc := os.Getenv("INPUT_BCC")
	replyTo := os.Getenv("INPUT_REPLY_TO")
	inReplyTo := os.Getenv("INPUT_IN_REPLY_TO")
	convertMarkdown := convertBool(os.Getenv("INPUT_CONVERT_MARKDOWN"))
	attachments := os.Getenv("INPUT_ATTACHMENTS")
	priority := os.Getenv("INPUT_PRIORITY")

	m := gomail.NewMessage()
	m.SetHeader("Subject", subject)
	m.SetHeader("From", m.FormatAddress(username, from))
	m.SetHeader("To", strings.Split(to, ",")...)
	if len(cc) != 0 {
		m.SetHeader("Cc", strings.Split(cc, ",")...)
	}
	if len(bcc) != 0 {
		m.SetHeader("Bcc", strings.Split(bcc, ",")...)
	}
	if len(replyTo) != 0 {
		m.SetHeader("Reply-To", replyTo)
	}
	if len(inReplyTo) != 0 {
		m.SetHeader("In-Reply-To", inReplyTo)
	}
	if len(priority) != 0 {
		m.SetHeader("X-Priority", priority)
	}

	if len(attachments) != 0 {
		for _, a := range strings.Split(attachments, ",") {
			m.Attach(a)
		}
	}

	if body != "" {
		m.SetBody("text/plain", getBody(body, false))
	}
	if htmlBody != "" {
		m.SetBody("text/html", getBody(htmlBody, convertMarkdown))
	}

	if mailPort == 0 {
		mailPort = 25
	}

	dialer := gomail.NewDialer(mailHost, mailPort, username, password)
	if dialer.SSL {
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := dialer.DialAndSend(m)
	if err != nil {
		log.Fatal(time.Now(), "send mail error: ", err)
	}
}

func getBody(bodyOrFile string, convertMarkdown bool) string {
	var body = bodyOrFile
	if strings.HasPrefix(bodyOrFile, "file://") {
		file := strings.TrimPrefix(bodyOrFile, "file://")
		bodyBts, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
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

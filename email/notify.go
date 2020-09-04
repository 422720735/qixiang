package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"qixiang/router"
	"strings"
	"time"
)

func getBody(city, user, to string) ([]byte, error) {
	res, err := router.Assemble(city)

	if err != nil {
		log.Printf("fetcher err %v", err)
		return nil, err
	}

	tmpl := template.New("index.html")
	tmpl = tmpl.Funcs(template.FuncMap{"unescaped": router.Unescaped})
	tmpl, err = tmpl.ParseFiles("index.html")

	if err != nil {
		log.Printf("templete paser err %v", err)
		return nil, err
	}

	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	subject := fmt.Sprintf("今日%s%s气温：%s%s",
		res.Pm25.CityName,
		res.CityLive.Today,
		res.CityLive.Situation,
		res.CityLive.Temperature,
	)

	body.Write([]byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + headers + "\r\n\r\n"))
	err = tmpl.Execute(&body, res)
	if err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}

func SendMail(city, user, password, host, to string) {
	body, err := getBody(city, user, to)
	if err != nil {
		log.Printf("fetcher err %v", err)
		return
	}
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	send_to := strings.Split(to, ";")
	err = smtp.SendMail(host, auth, user, send_to, body)
	if err != nil {
		fmt.Printf("%s 发送邮件失败，err：%v", time.Now().Format("2006-01-02 15:04:05"), err)
	} else {
		fmt.Printf("%s 发送邮件成功\n", time.Now().Format("2006-01-02 15:04:05"))
	}
}

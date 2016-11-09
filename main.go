package main

import (
	"fmt"

	"io/ioutil"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	//	file, err := os.Open(cnf.Mail.Body) // For read access.

	buff, err := ioutil.ReadFile(cnf.Mail.Body)
	if err != nil {
		fmt.Println(err)
	}
	body := string(buff)
	sendEmail(body, cnf.Mail.Attach)
	fmt.Println("end")

}

func sendEmail(Body string, filename string) {
	fmt.Println("Sending email")
	for _, to := range cnf.Mail.To {
		mail := gomail.NewMessage()
		mail.SetHeader("From", cnf.Mail.From)
		mail.SetHeader("To", to)
		mail.SetHeader("Subject", cnf.Mail.Subject)
		mail.SetBody("text/html", Body)
		mail.Attach(filename)
		d := gomail.Dialer{Host: cnf.Mail.Host, Port: cnf.Mail.Port}
		if err := d.DialAndSend(mail); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Email sent!")
	}

}

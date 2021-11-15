package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"smail/util"

	"crypto/tls"

	"gopkg.in/gomail.v2"
	// Если надо парсить
	//"strings"
)

func main() {
	toMail := flag.String("to", "", "input path")
	subjMail := flag.String("subj", "", "output path")
	mailBody := flag.String("data", "", "Mail Message")
	mailAttach := flag.String("att", "", "Attach file")
	flag.Parse()

	// Парсим строку с адресами электронной почты вида "info@gmail.com;result@gmail.com"
	// var toMail1 = strings.Split(*toMail, ";")
	// LoadConfig reads configuration from file or environment variables.

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	result, err := ioutil.ReadFile(*mailBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(config.Server)
	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	// m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	// Если нужно несколько получателей
	//m.SetHeader("To", toMail1...)
	m.SetHeader("To", *toMail)
	m.SetHeader("Subject", *subjMail)
	m.SetBody("text/html", string(result))
	if *mailAttach != "" {
		m.Attach(*mailAttach)
	}

	d := gomail.NewDialer(config.Server, config.Port, config.Login, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// d.TLSConfig = &tls.Config{ServerName: "MOS160.moscow.eurochem.ru", InsecureSkipVerify: false}

	if err := d.DialAndSend(m); err != nil {
		// panic(err)
		fmt.Println(err)
	}

}

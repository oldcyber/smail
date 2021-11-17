package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"smail/util"
	"strings"

	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func Split(r rune) bool {
	return r == ';'
}

func main() {
	toMail := flag.String("to", "", "input path (split with ';')")
	subjMail := flag.String("subj", "", "output path")
	mailBody := flag.String("data", "", "Mail Message")
	mailAttach := flag.String("att", "", "Attach file (split with ';')")
	flag.Parse()

	toMailValue := *toMail
	mailAttachValue := *mailAttach
	toMailArr := strings.FieldsFunc(toMailValue, Split)
	mailAttachArr := strings.FieldsFunc(mailAttachValue, Split)

	for i := 0; i < len(toMailArr); i++ {

	}

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	result, err := ioutil.ReadFile(*mailBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	d := gomail.NewDialer(config.Server, config.Port, config.Login, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	for _, r := range toMailArr {
		m.SetHeader("From", config.From)
		m.SetHeader("To", r)
		m.SetHeader("Subject", *subjMail)
		m.SetBody("text/html", string(result))
		if len(mailAttachArr) > 0 {
			for _, i := range mailAttachArr {
				m.Attach(i)
			}
			// for i := 0; i < len(mailAttachArr); i++ {
			// 	m.Attach(mailAttachArr[i])
			// }

		}
		if err := d.DialAndSend(m); err != nil {
			// panic(err)
			fmt.Println(err)
		}

		m.Reset()
	}

}

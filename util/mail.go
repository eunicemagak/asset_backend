package util

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/smtp"
// 	"strings"
// )

// const SENDER = "victorkbuzes@gmail.com"

// type Mail struct {
// 	Sender  string
// 	To      []string
// 	Cc      []string
// 	Bcc     []string
// 	Subject string
// 	Body    string
// }

// type SmtpServer struct {
// 	Host      string
// 	Port      string
// 	TlsConfig *tls.Config
// }

// func (s *SmtpServer) ServerName() string {
// 	return s.Host + ":" + s.Port
// }

// func (mail *Mail) BuildMessage() []byte {
// 	mail.Sender = SENDER

// 	header := ""
// 	header += fmt.Sprintf("From:%s\r\n", mail.Sender)
// 	if len(mail.To) > 0 {
// 		header += fmt.Sprintf("To:%s\r\n", strings.Join(mail.To, ";"))
// 	}
// 	if len(mail.Cc) > 0 {
// 		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
// 	}

// 	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
// 	header += "\r\n" + mail.Body

// 	return []byte(header)

// }

// func (mail *Mail) SendMail(messageBody []byte) {
// 	smtpServer := SmtpServer{Host: "smtp.gmail.com", Port: "465"}
// 	smtpServer.TlsConfig = &tls.Config{
// 		InsecureSkipVerify: true,
// 		ServerName:         smtpServer.Host,
// 	}

// 	auth := smtp.PlainAuth("", mail.Sender, "33947267", smtpServer.Host)

// 	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TlsConfig)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	client, err := smtp.NewClient(conn, smtpServer.Host)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer client.Quit()
// 	// step 1: Use Auth
// 	if err = client.Auth(auth); err != nil {
// 		log.Panic(err)
// 	}

// 	// step 2: add all from and to
// 	if err = client.Mail(mail.Sender); err != nil {
// 		log.Panic(err)
// 	}
// 	receivers := append(mail.To, mail.Cc...)
// 	receivers = append(receivers, mail.Bcc...)
// 	for _, k := range receivers {
// 		log.Println("sending to: ", k)
// 		if err = client.Rcpt(k); err != nil {
// 			log.Panic(err)
// 		}
// 	}

// 	// Data
// 	w, err := client.Data()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	_, err = w.Write(messageBody)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = w.Close()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	log.Println("Mail sent successfully")
// }

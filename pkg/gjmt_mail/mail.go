package gjmt_mail

import (
	"gopkg.in/gomail.v2"
	"os"
)

// Config хранит конфигурацию для SMTP
type Config struct {
	Mailer     string
	Host       string
	Port       string
	Encryption string
	FromName   string
	Username   string
	Password   string
}

// NewConfig создает новый конфиг из переменных окружения
func NewConfig() *Config {
	return &Config{
		Mailer:     os.Getenv("MAIL_MAILER"),
		Host:       os.Getenv("MAIL_HOST"),
		Port:       os.Getenv("MAIL_PORT"),
		Encryption: os.Getenv("MAIL_ENCRYPTION"),
		FromName:   os.Getenv("MAIL_FROM_NAME"),
		Username:   os.Getenv("MAIL_USERNAME"),
		Password:   os.Getenv("MAIL_PASSWORD"),
	}
}

// SendMail отправляет письмо по SMTP
func (c *Config) SendMail(to string, subject, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", c.Username)
	m.SetHeader("From", os.Getenv("MAIL_FROM_NAME")+" <"+c.Username+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject+" — "+os.Getenv("APP_NAME"))
	m.SetBody("text/html", body)
	//	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(c.Host, 2525, c.Username, c.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}

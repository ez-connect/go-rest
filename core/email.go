package core

import (
	"errors"
	"fmt"
	"net/smtp"
	"sync"
)

// Use login auth instead of PlainAuth
// that was support by golang
type Auth struct {
	UserName string
	Password string
}

func (a *Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.UserName), nil
}

func (a *Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.UserName), nil
		case "Password:":
			return []byte(a.Password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

///////////////////////////////////////////////////////////////////

// Config for SMTP
type SMTPConfig struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
}

type SMTP struct {
	config SMTPConfig
}

var smtpOnce sync.Once
var smptAuth smtp.Auth
var smtpServer *SMTP

///////////////////////////////////////////////////////////////////

// Singleton pattern
func GetSMTP() *SMTP {
	smtpOnce.Do(func() {
		smtpServer = &SMTP{}
	})

	return smtpServer
}

///////////////////////////////////////////////////////////////////

func (s *SMTP) Init(config SMTPConfig) {
	s.config = config
}

func (s *SMTP) LoginAuth(username, password string) smtp.Auth {
	return &Auth{
		UserName: username,
		Password: password,
	}
}

// Connect to the server, authenticate, set the sender and recipient,
// and send the email all in one step
func (s *SMTP) Send(from string, subject, body string, to string, toAddressess []string) error {
	addr := fmt.Sprintf("%s:%v", s.config.Server, s.config.Port)
	// Set up authentication information
	smptAuth = s.LoginAuth(
		s.config.Username,
		s.config.Password,
	)

	msg := fmt.Sprintf(
		"To: %s\nSubject: %s\nContent-Type: text/html;\n\n%s",
		to, subject, body,
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step
	return smtp.SendMail(
		addr,
		smptAuth,
		from,
		toAddressess,
		[]byte(msg),
	)
}

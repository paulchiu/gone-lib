package email

import (
	"net/smtp"
	"regexp"
)

type Server struct {
	Username string
	Password string
	Host string
}

// Send plain text email
func SendPlain(server *Server, message *Message) error {
	return send(server, message, message.Plain())
}

// Send html email
func SendHTML(server *Server, message *Message) error {
	return send(server, message, message.HTML())
}

// Generic send
func send(server *Server, message *Message, messageContent []byte) error {
	auth := smtp.PlainAuth("", server.Username, server.Password, stripPort(server.Host))
	err := smtp.SendMail(server.Host, auth, message.From.Address, addressListToStrings(message.To), messageContent)
	return err
}

// Remove port from host string
func stripPort(host string) string {
	portPattern := regexp.MustCompile(`:\d+`)
	return portPattern.ReplaceAllString(host, "")
}

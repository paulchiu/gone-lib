package email

import (
	"net/smtp"
	"regexp"
)

// Email server connection details
type Connection struct {
	Username string
	Password string
	Host     string
}

// Send plain text email using a given connection
func (conn *Connection) SendPlain(message *Message) error {
	return conn.send(message, message.Plain())
}

// Send html email using a given connection
func (conn *Connection) SendHTML(message *Message) error {
	return conn.send(message, message.HTML())
}

// Generic send
func (conn *Connection) send(message *Message, messageContent []byte) error {
	auth := smtp.PlainAuth("", conn.Username, conn.Password, stripPort(conn.Host))
	err := smtp.SendMail(conn.Host, auth, message.From.Address, addressListToStrings(message.To), messageContent)
	return err
}

// Remove port from host string
func stripPort(host string) string {
	portPattern := regexp.MustCompile(`:\d+`)
	return portPattern.ReplaceAllString(host, "")
}

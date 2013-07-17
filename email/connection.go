package email

import (
	"net/smtp"
	"regexp"
)

// Email server connection details.
type Connection struct {
	Username string
	Password string
	Host     string
}

// Send byte message using format.
func (conn *Connection) Send(message *Message, format string) error {
	auth := smtp.PlainAuth("", conn.Username, conn.Password, stripPort(conn.Host))
	err := smtp.SendMail(conn.Host, auth, message.From.Address, message.To.ToStrings(), message.ToBytes(format))
	return err
}

// Remove port from host string.
func stripPort(host string) string {
	portPattern := regexp.MustCompile(`:\d+`)
	return portPattern.ReplaceAllString(host, "")
}

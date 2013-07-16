package email

import (
	"net/mail"
	"net/smtp"
	"regexp"
)

const CRLF = "\r\n"

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

// Convert an array of RFC 5322 addresses to strings
func addressListToStrings(addressList []*mail.Address) []string {
	return addressListTo(addressList, func (a *mail.Address) string { return a.Address })
}

// Convert an array of RFC 5322 addresses to RFC 5322 strings
func addressListToRFCStrings(addressList []*mail.Address) []string {
	return addressListTo(addressList, func (a *mail.Address) string { return a.String() })
}

// Covert an array of RFC 5322 to a given function utilising an address
func addressListTo(addressList []*mail.Address, f func (a *mail.Address) string) []string {
	strings := make([]string, len(addressList))
	for index, address := range addressList {
		strings[index] = f(address)
	}
	return strings
}

// Remove port from host string
func stripPort(host string) string {
	portPattern := regexp.MustCompile(`:\d+`)
	return portPattern.ReplaceAllString(host, "")
}

package email

import (
	"net/mail"
	"net/smtp"
	"regexp"
	"errors"
	"strings"
)

const CRLF = "\r\n"

type Server struct {
	Username string
	Password string
	Host string
}

type Message struct {
	From *mail.Address
	To []*mail.Address
	subject string
	body string
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

// Create a new email message
func NewMessage(from, to, subject, body string) *Message {
	message, _ := NewMessageDebug(from, to, subject, body)
	return message
}

// Create a new email message, return any reasons for failure
func NewMessageDebug(from, to, subject, body string) (*Message, error) {
	fromAddress, err := mail.ParseAddress(from)
	if err != nil {
		return nil, errors.New("Invalid from address; " + err.Error())
	}

	toAddresses, err := mail.ParseAddressList(to)
	if err != nil {
		return nil, errors.New("Invalid to address(es); " + err.Error())
	}

	return &Message{fromAddress, toAddresses, subject, body}, nil
}

// Generate default message header
func (message *Message) header() string {
	header := "From: " + message.From.String() + CRLF
	header += "To: " + strings.Join(addressListToRFCStrings(message.To), ", ") + CRLF
	header += "Subject: " + message.subject + CRLF
	return header
}

// Generate html message header
func (message *Message) htmlHeader() string {
	return message.header() + "Content-Type: text/html" + CRLF
}

// Convert email message to plain format string
func (message *Message) PlainString() string {
	return message.header() + CRLF + message.body
}

// Convert email message to plain format string
func (message *Message) HTMLString() string {
	return message.htmlHeader() + CRLF + message.body
}

// Convert email message to plain format byte array for sending
func (message *Message) Plain() []byte {
	return []byte(message.PlainString())
}

// Convert email message to html format byte array for sending
func (message *Message) HTML() []byte {
	return []byte(message.HTMLString())
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

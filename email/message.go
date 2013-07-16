package email

import (
	"net/mail"
	"strings"
	"errors"
)

type Message struct {
	From *mail.Address
	To []*mail.Address
	subject string
	body string
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


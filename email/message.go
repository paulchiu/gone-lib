package email

import (
	"errors"
	"net/mail"
	"strings"
)

// Message header separator.
const CRLF = "\r\n"

// Email message.
type Message struct {
	From    *mail.Address
	To      *AddressList
	Subject string
	Body    string
}

// Create a new email message, assuming no errors.
func NewMessage(from, to, subject, body string) *Message {
	message, _ := NewMessageDebug(from, to, subject, body)
	return message
}

// Create a new email message, returning any reasons for failure.
func NewMessageDebug(from, to, subject, body string) (*Message, error) {
	fromAddress, err := mail.ParseAddress(from)
	if err != nil {
		return nil, errors.New("Invalid from address; " + err.Error())
	}

	toAddresses, err := mail.ParseAddressList(to)
	if err != nil {
		return nil, errors.New("Invalid to address(es); " + err.Error())
	}

	return &Message{fromAddress, &AddressList{toAddresses}, subject, body}, nil
}

// Generate plain text message header.
func (message *Message) plainHeader() string {
	header := "From: " + message.From.String() + CRLF
	header += "To: " + strings.Join(message.To.ToRFCStrings(), ", ") + CRLF
	header += "Subject: " + message.Subject + CRLF
	return header
}

// Generate html message header.
func (message *Message) htmlHeader() string {
	return message.plainHeader() + "Content-Type: text/html" + CRLF
}

// Convert email message to plain text message format, return as string.
func (message *Message) PlainString() string {
	return message.plainHeader() + CRLF + message.Body
}

// Convert email message to html message format, return as string.
func (message *Message) HTMLString() string {
	return message.htmlHeader() + CRLF + message.Body
}

// Convert email message to a given format in byte form, which is needed by stmp.SendMail(). Supports either "plain" or "html".
func (message *Message) ToBytes(format string) []byte {
	var byteData []byte

	switch format {
	case "plain":
		byteData = []byte(message.PlainString())
	case "html":
		byteData = []byte(message.HTMLString())
	default:
		panic("Unsupported message format")
	}

	return byteData
}

// Copyright 2013 Paul Chiu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package email

import (
	"errors"
	"net/mail"
	"strings"
)

const CRLF = "\r\n"

type Message struct {
	From    *mail.Address
	To      []*mail.Address
	subject string
	body    string
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

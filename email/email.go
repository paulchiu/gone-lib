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
	"net/smtp"
	"regexp"
)

type Server struct {
	Username string
	Password string
	Host     string
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

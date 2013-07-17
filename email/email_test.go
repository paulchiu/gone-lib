package email

import (
	"testing"
)

var plainFixture = map[string]string{
	"host":     "smtp.gmail.com:587",
	"username": "username@gmail.com",
	"password": "password",
	"from":     "username@gmail.com",
	"to":       "example@gmail.com",
	"subject":  "hello world",
	"body":     "lorem ipsum etc.\n\nregards,",
}

var htmlFixture = map[string]string{
	"host":     plainFixture["host"],
	"username": plainFixture["username"],
	"password": plainFixture["password"],
	"from":     "Mr From <" + plainFixture["from"] + ">",
	"to":       "Ms To <" + plainFixture["to"] + ">",
	"subject":  "hello html world",
	"body":     "<h1>lorem ipsum</h1><p>something something etc.</p><p>regards,</p>",
}

func TestSend(t *testing.T) {
	err := Send(htmlFixture["host"], htmlFixture["username"], htmlFixture["password"], htmlFixture["from"], htmlFixture["to"], htmlFixture["subject"], htmlFixture["body"])

	if err != nil {
		t.Error("Could not send html message; " + err.Error())
	}
}

func TestSendPlain(t *testing.T) {
	err := SendPlain(plainFixture["host"], plainFixture["username"], plainFixture["password"], plainFixture["from"], plainFixture["to"], plainFixture["subject"], plainFixture["body"])

	if err != nil {
		t.Error("Could not send plain message; " + err.Error())
	}
}

// This example sends a HTML email using "one" line of code using Gmail's SMTP server.
func ExampleSend() {
	err := Send("smtp.gmail.com:587",
		"username@gmail.com",
		"password",
		"Abe <username@gmail.com>",
		"bob@gmail.com",
		"Hello Bob",
		"<p>Hi Bob</p><p>Long time no see.</p><p>Abe</p>")

	if err != nil {
		// Could not send email
	}
}

// This example sends a plain email using "one" line of code using a random SMTP server.
func ExampleSendPlain() {
	err := SendPlain("smtp.example.com:25",
		"username@example.com",
		"password",
		"Abe <username@example.com>",
		"bob@example.com",
		"Hello Bob",
		"Hi Bob\n\nLong time no see.\n\nAbe")

	if err != nil {
		// Could not send email
	}
}

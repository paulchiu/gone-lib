package email

import (
	"testing"
)

var conn = &Connection{"username@gmail.com", "password", "smtp.gmail.com:587"}

var plainFixture = map[string]string{
	"from":    "from@gmail.com",
	"to":      "to@gmail.com",
	"subject": "hello world",
	"body":    "lorem ipsum etc.\n\nregards,",
}

var htmlFixture = map[string]string{
	"from":    "Mr From <" + plainFixture["from"] + ">",
	"to":      "Ms To <" + plainFixture["to"] + ">",
	"subject": "hello html world",
	"body":    "<h1>lorem ipsum</h1><p>something something etc.</p><p>regards,</p>",
}

func TestStripPort(t *testing.T) {
	host := "smtp.gmail.com:587"
	hostServer := stripPort(host)

	if hostServer != "smtp.gmail.com" {
		t.Error("Strip port did not work")
	}
}

func ignoreTestPlain(t *testing.T) {
	message := NewMessage(plainFixture["from"], plainFixture["to"], plainFixture["subject"], plainFixture["body"])
	err := conn.SendPlain(message)

	if err != nil {
		t.Error("Could not send plain message; " + err.Error())
	}
}

func ignoreTestHTML(t *testing.T) {
	message := NewMessage(htmlFixture["from"], htmlFixture["to"], htmlFixture["subject"], htmlFixture["body"])
	err := conn.SendHTML(message)

	if err != nil {
		t.Error("Could not send html message; " + err.Error())
	}
}

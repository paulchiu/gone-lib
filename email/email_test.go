package email

import (
	"net/mail"
	"testing"
)

var server = &Server{"username@gmail.com", "password", "smtp.gmail.com:587"}

func TestAddressConvertToString(t *testing.T) {
	addressesB, _ := mail.ParseAddressList("hello@example.com, world@example.com")
	strAddresses := addressListToStrings(addressesB)

	if len(strAddresses) != 2 {
		t.Error("Unexpected post-conversion length")
	}

	if strAddresses[0] != "hello@example.com" {
		t.Error("First address does not match")
	}

	if strAddresses[1] != "world@example.com" {
		t.Error("Second address does not match")
	}
}

func TestAddressConvertToRFC(t *testing.T) {
	addressesB, _ := mail.ParseAddressList("hello@example.com, world@example.com")
	strAddresses := addressListToRFCStrings(addressesB)

	if len(strAddresses) != 2 {
		t.Error("Unexpected post-conversion length")
	}

	if strAddresses[0] != "<hello@example.com>" {
		t.Error("First address does not match")
	}

	if strAddresses[1] != "<world@example.com>" {
		t.Error("Second address does not match")
	}
}

func TestStripPort(t *testing.T) {
	host := "smtp.gmail.com:587"
	hostServer := stripPort(host)

	if hostServer != "smtp.gmail.com" {
		t.Error("Strip port did not work")
	}
}

func TestPlainMessage(t *testing.T) {
	message, err := NewMessageDebug("from@gmail.com", "to@gmail.com", "hello world", "lorem ipsum etc.\n\nregards,")

	if err != nil {
		t.Error(err)
	} else {
		t.Log(message.PlainString())
	}
}

func TestHTMLMessage(t *testing.T) {
	message, err := NewMessageDebug("Mr From <from@gmail.com>", "Ms To <to@gmail.com>", "hello html world", "<h1>lorem ipsum</h1><p>something something etc.</p><p>regards,</p>")

	if err != nil {
		t.Error(err)
	} else {
		t.Log(message.HTMLString())
	}
}

func xTestPlain(t *testing.T) {
	message := NewMessage("from@gmail.com", "to@gmail.com", "hello world", "lorem ipsum etc.\n\nregards,")
	err := SendPlain(server, message)

	if err != nil {
		t.Error("Could not send plain message; " + err.Error())
	}
}

func xTestHTML(t *testing.T) {
	message := NewMessage("Mr From <from@gmail.com>", "Ms To <to@gmail.com>", "hello html world", "<h1>lorem ipsum</h1><p>something something etc.</p><p>regards,</p>")
	err := SendHTML(server, message)

	if err != nil {
		t.Error("Could not send html message; " + err.Error())
	}
}

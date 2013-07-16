package email

import (
	"testing"
)

func TestPlainMessage(t *testing.T) {
	message, err := NewMessageDebug(plainFixture["from"], plainFixture["to"], plainFixture["subject"], plainFixture["body"])

	if err != nil {
		t.Error(err)
	} else {
		t.Log(message.PlainString())
	}
}

func TestHTMLMessage(t *testing.T) {
	message, err := NewMessageDebug(htmlFixture["from"], htmlFixture["to"], htmlFixture["subject"], htmlFixture["body"])

	if err != nil {
		t.Error(err)
	} else {
		t.Log(message.HTMLString())
	}
}

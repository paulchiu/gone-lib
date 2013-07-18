package email

import (
	"bytes"
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

func TestToBytes(t *testing.T) {
	message := NewMessage(htmlFixture["from"], htmlFixture["to"], htmlFixture["subject"], htmlFixture["body"])

	if bytes.Compare([]byte(message.HTMLString()), message.ToBytes("html")) != 0 {
		t.Error("ToBytes(html) failed")
	}
	
	if bytes.Compare([]byte(message.PlainString()), message.ToBytes("plain")) != 0 {
		t.Error("ToBytes(plain) failed")
	}
}

func TestInvalidToBytes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Did not panic when attempting to convert message to invalid byte format")
		}
	}()

	message := NewMessage(htmlFixture["from"], htmlFixture["to"], htmlFixture["subject"], htmlFixture["body"])
	message.ToBytes("invalid format")
}

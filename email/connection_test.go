package email

import (
	"testing"
)

func TestStripPort(t *testing.T) {
	host := "smtp.gmail.com:587"
	hostServer := stripPort(host)

	if hostServer != "smtp.gmail.com" {
		t.Error("Strip port did not work")
	}
}

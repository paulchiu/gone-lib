package email

import (
	"net/mail"
	"testing"
)

var sampleAddressList = "hello@example.com, Earth <world@example.com>"

func TestAddressConvertToStrings(t *testing.T) {
	addresses, _ := mail.ParseAddressList(sampleAddressList)
	addressList := &AddressList{addresses}
	strAddresses := addressList.ToStrings()

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
	addresses, _ := mail.ParseAddressList(sampleAddressList)
	addressList := &AddressList{addresses}
	strAddresses := addressList.ToRFCStrings()

	if len(strAddresses) != 2 {
		t.Error("Unexpected post-conversion length")
	}

	if strAddresses[0] != "<hello@example.com>" {
		t.Error("First address does not match")
	}

	if strAddresses[1] != "\"Earth\" <world@example.com>" {
		t.Error("Second address does not match")
	}
}

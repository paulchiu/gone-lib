package email

import (
	"net/mail"
)

// AddressList type; essentially just an alias for array of Addresses.
type AddressList struct {
	items []*mail.Address
}

// Convert an array of RFC 5322 addresses to strings usable as destination emails.
func (addressList *AddressList) ToStrings() []string {
	return addressList.transform(func(a *mail.Address) string { return a.Address })
}

// Convert an array of RFC 5322 addresses to RFC 5322 strings; i.e. "Name <email@example.com>".
func (addressList *AddressList) ToRFCStrings() []string {
	return addressList.transform(func(a *mail.Address) string { return a.String() })
}

// Transform an array of RFC 5322 addresses to a string using a given function.
func (addressList *AddressList) transform(f func(a *mail.Address) string) []string {
	strings := make([]string, len(addressList.items))
	for index, address := range addressList.items {
		strings[index] = f(address)
	}
	return strings
}

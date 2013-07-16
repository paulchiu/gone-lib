package email

import (
	"net/mail"
)

// Convert an array of RFC 5322 addresses to strings
func addressListToStrings(addressList []*mail.Address) []string {
	return addressListTo(addressList, func(a *mail.Address) string { return a.Address })
}

// Convert an array of RFC 5322 addresses to RFC 5322 strings
func addressListToRFCStrings(addressList []*mail.Address) []string {
	return addressListTo(addressList, func(a *mail.Address) string { return a.String() })
}

// Covert an array of RFC 5322 to a given function utilising an address
func addressListTo(addressList []*mail.Address, f func(a *mail.Address) string) []string {
	strings := make([]string, len(addressList))
	for index, address := range addressList {
		strings[index] = f(address)
	}
	return strings
}

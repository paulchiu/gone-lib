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

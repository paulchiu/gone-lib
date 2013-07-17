// Gone-lib email "one" liners. See func Send() for how to quickly send an email.
package email

// Send new email message. Assumed to be HTML message.
//
// Parameters:
//		* host - SMTP server, should include port; i.e. "smtp.example.com:25".
//		* username - SMTP server username.
//		* password - SMTP server password.
//		* from - From email; i.e. "alice@example.com" or "Alice Bane <alice@example.com>".
//		* to - To email(s); i.e. "bob@example.com" or "Bob Cane <bob@example.com>, Cat Doe <cat@xample.com>".
//		* subject - Subject for the email.
//		* body - HTML formatted email body; i.e. "<p>Hello</p>".
func Send(host, username, password, from, to, subject, body string) error {
	return send(host, username, password, from, to, subject, body, "html")
}

// Send new plain email message. Same parameters as Send(), however, body is expected to be plain text.
func SendPlain(host, username, password, from, to, subject, body string) error {
	return send(host, username, password, from, to, subject, body, "plain")
}

// Send new email message with given format.
func send(host, username, password, from, to, subject, body, format string) error {
	message, err := NewMessageDebug(from, to, subject, body)
	if err != nil {
		return err
	}

	conn := &Connection{username, password, host}

	return conn.Send(message, format)
}

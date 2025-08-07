package user

import "net/mail"

type Email string

func NewEmail(s string) (Email, error) {
	email, err := mail.ParseAddress(s)
	if err != nil {
		return "", err
	}

	return Email(email.Address), nil
}

package user

import (
	"errors"
	"unicode"
)

type Password string

func NewPassword(s string) (Password, []error) {
	var (
		upper, special, digit = false, false, false
	)

	errs := make([]error, 0, 4)

	for _, c := range s {
		if unicode.IsDigit(c) {
			digit = true
		} else if unicode.IsUpper(c) {
			upper = true
		} else if c >= '!' && c <= '/' {
			special = true
		} else if c >= ':' && c <= '@' {
			special = true
		} else if c >= '[' && c <= '`' {
			special = true
		} else if c >= '{' && c <= '~' {
			special = true
		}
	}

	if !upper {
		errs = append(errs, errors.New("there's no uppercase char"))
	}
	if !special {
		errs = append(errs, errors.New("there's no special char"))
	}
	if !digit {
		errs = append(errs, errors.New("there's no numeric char"))
	}
	if len(s) < 8 {
		errs = append(errs, errors.New("password too short"))
	}

	if len(errs) > 0 {
		return "", errs
	}

	return Password(s), nil
}

package user

import (
	"errors"
	"strings"
)

type UserName string

var (
	ErrNameTooShort = errors.New("name needs at least 3 chars")
	ErrNameTooLong  = errors.New("name should contain a maximum of 45 chars")
	ErrEmptyName    = errors.New("name should contain between 3 to 45 chars")
)

func NewUserName(s string) (UserName, error) {
	if len(s) < 3 {
		return "", ErrNameTooShort
	}

	if len(s) > 45 {
		return "", ErrNameTooLong
	}

	if s := strings.TrimSpace(s); len(s) == 0 {
		return "", ErrEmptyName
	}

	return UserName(s), nil
}

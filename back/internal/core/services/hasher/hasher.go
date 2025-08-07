package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func (h *Hasher) Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 13)
	return string(bytes), err
}

func (h *Hasher) CompareHashAndPasswd(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

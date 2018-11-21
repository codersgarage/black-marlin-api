package utils

import "golang.org/x/crypto/bcrypt"

func CreateHashedPassword(v string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(pass)
}

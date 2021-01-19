package middleware

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	cost := 8 //Cost is the number of times you will encrypt the password
	bytes, err := bcrypt.GenerateFromPassword([]bytes(pass), cost)
	return string(bytes), err
}

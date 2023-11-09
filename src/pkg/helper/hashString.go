package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(str string) (string, error) {

	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedString), nil
}

func CompareHashWithHashString(str string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(hashed))

	return err != nil
}

package helper

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RandomCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

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

func GetValueType(args ...interface{}) string {

	for _, v := range args {

		switch v.(type) {

		case int:
			return "int"
		case int64:
			return "int"
		case int32:
			return "int"
		case string:
			return "string"
		case uint:
			return "int"

		default:

			return ""

		}

	}
	return ""

}

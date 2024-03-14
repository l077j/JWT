package jwt

import "github.com/dgrijalva/jwt-go"

type Data struct {
	Name   string
	Age    int
	Gender string
	jwt.StandardClaims
}

var key = "abcdefg1234567"

func Sign(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	sign, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return sign, nil
}

func Verify(sign string, data jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return err
}

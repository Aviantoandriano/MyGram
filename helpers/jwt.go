package helpers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web/context"

	jwt "github.com/dgrijalva/jwt-go"
)

var SecretKey *ecdsa.PrivateKey

func init() {
	JwtKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	SecretKey = JwtKey
}

func GenerateToken(id int, email string) (token string) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := parseToken.SignedString(SecretKey)
	if err != nil {
		panic(err)
	}
	return
}

func VerifyToken(c *context.Context) (response interface{}, err error) {
	err = errors.New("sign in proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		response = nil
		return
	}
	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			response = nil
			return response, err
		}
		return SecretKey, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return
	}
	response = token.Claims.(jwt.MapClaims)
	err = nil
	return
}

func ValidateToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}
		return SecretKey, nil
	})
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return token, err
	}
	return token, err
}

package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go-rest-api/src/constant"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(constant.SampleSecretKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(bearerToken string) (err error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return constant.SampleSecretKey, nil
	})
	if err != nil {
		return err
	}

	if token == nil {
		err = errors.New("token error")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("invalid JWT Token")
		return err
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return nil
}

func GenerateJWT2(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = username
	claims["exp"] = time.Now().Add(10 * time.Minute)

	tokenString, err := token.SignedString(constant.SampleSecretKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {})
}

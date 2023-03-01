package utlis

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/golang-jwt/jwt"
	"go-rest-api/src/constant"
)

func generateJWT2(username string) (string, error) {
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

func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {})
}
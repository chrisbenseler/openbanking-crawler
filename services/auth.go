package services

import (
	"fmt"
	"net/http"
	"openbankingcrawler/common"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret = "SAMPLE"
)

//Auth service interface
type Auth interface {
	ValidateAccessToken(*http.Request) (*jwt.Token, common.CustomError)
	CreateAccessToken(email string, password string) (*string, common.CustomError)
}

type auth struct {
}

//NewAuthService create a new auth service
func NewAuthService() Auth {
	return &auth{}
}

//ValidateAccessToken authenticate a request
func (s *auth) ValidateAccessToken(request *http.Request) (*jwt.Token, common.CustomError) {

	tokenString := extractToken(request)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, common.NewUnauthorizedError("Bad credentials")
	}
	return token, nil
}

//CreateAccessToken create an access token
func (s *auth) CreateAccessToken(email string, password string) (*string, common.CustomError) {

	if email != "bvlab@bv.com.br" || password != "abcd1234" {
		return nil, common.NewUnauthorizedError("Bad credentials")
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))

	if err != nil {
		return nil, common.NewInternalServerError("Token error", err)
	}
	return &token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

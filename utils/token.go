package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaasikodes/e-commerce-go/constants"
	"github.com/kaasikodes/e-commerce-go/models"
)

func CreateJWT( userId string) (string, error) {
	secret := constants.JWTSecret

	expiration := constants.JWTExpirationTime
	userIdKey := constants.JWTUserIdMapKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(userIdKey): userId,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err
}

func GetAccessTokenFromRequest(r *http.Request) (string, error) {
	tokenHeader := r.Header.Get("Authorization")
	// The usual convention is for "Bearer" to be title-cased. However, there's no
	// strict rule around this, and it's best to follow the robustness principle here.
	if len(tokenHeader) < 7 || !strings.EqualFold(tokenHeader[:7], "bearer ") {
		return "", fmt.Errorf("no token present in request header")
	}
	return tokenHeader[7:], nil
}

func ValidateJWT (tokenStr string) (*jwt.Token , error) {
	secret := constants.JWTSecret
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}



func RetrieveUserFromRequestContext(r *http.Request) (models.User, error){
	userKey := constants.JWTAuthUserContextKey
	user, ok := r.Context().Value(userKey).(models.User)

	if !ok {
		return models.User{}, fmt.Errorf("unable to retrieve user from context")
	}
	return user, nil
}
package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateTokens(session, uuid string) (refreshToken, accessToken string, err error) {
	accessToken = CreateAccessToken(session, uuid)
	refreshToken = CreateRefreshToken(session, uuid)
	return accessToken, refreshToken, err
}

func CreateAccessToken(session, uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"aud": []string{uuid},
		"jti": []string{session},
	})

	s, _ := t.SignedString([]byte(os.Getenv("SECRET")))
	return s
}

func CreateRefreshToken(session, uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"aud": []string{uuid},
		"jti": []string{session},
	})
	s, _ := t.SignedString([]byte(os.Getenv("SECRET")))
	return s
}

func ExtractClaims(tokenStr string) (models.JwtClaims, error) {
	claims := models.JwtClaims{}
	hmacSecretString := os.Getenv("SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return models.JwtClaims{}, err
	}

	claimsMap, _ := token.Claims.(jwt.MapClaims)

	expiresAt, _ := claimsMap["exp"]

	audience, _ := claimsMap["aud"].([]interface{})

	issuedAt, _ := claimsMap["iat"]

	session, _ := claimsMap["jti"].([]interface{})

	claims = models.JwtClaims{
		Audience:  audience[0].(string),
		ExpiresAt: expiresAt.(float64),
		IssuedAt:  issuedAt.(float64),
		Session:   session[0].(string),
	}
	return claims, nil

}

func GetTokenFromHeader(c *gin.Context) string {
	rawToken := c.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(rawToken, "Bearer ")
	return accessToken
}

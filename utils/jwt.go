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

func GenerateTokens(uuid string) (refreshToken, accessToken string, err error) {
	accessToken = CreateAccessToken(uuid)
	refreshToken = CreateRefreshToken(uuid)
	return accessToken, refreshToken, err
}

func CreateAccessToken(uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"aud": []string{uuid},
	})

	s, _ := t.SignedString([]byte(os.Getenv("SECRET")))
	return s
}

func CreateRefreshToken(uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"aud": []string{uuid},
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

	claims = models.JwtClaims{
		Audience:  audience[0].(string),
		ExpiresAt: expiresAt.(float64),
		IssuedAt:  issuedAt.(float64),
	}
	return claims, nil

}

func GetTokenFromHeader(c *gin.Context) string {
	rawToken := c.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(rawToken, "Bearer ")
	return accessToken
}

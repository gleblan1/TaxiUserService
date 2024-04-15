package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"reflect"
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
	return refreshToken, accessToken, err
}

func CreateAccessToken(uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"aud": []string{uuid},
	})
	s, _ := t.SignedString([]byte(os.Getenv("SECRET")))
	return s
}

//нужно проверять передаваемый пользователем в запросе токен, сравнивая его с токеном из бд. в бд будет вайтлист, в котором будут храниться пары. если токен в вайтлисте то все ок.
//если пользователь лог аут то из вайтлиста удаляется пара (?)

func CreateRefreshToken(uuid string) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30 * time.Duration(1)).Unix(),
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
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return models.JwtClaims{}, err
	}

	claimsMap, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.JwtClaims{}, fmt.Errorf("Token claims are not of type jwt.MapClaims")
	}

	expiresAt, ok := claimsMap["exp"]
	if !ok {
		return models.JwtClaims{}, fmt.Errorf("Claim 'exp' is not of type int64")
	}

	audience, ok := claimsMap["aud"].([]interface{})
	if !ok {
		return models.JwtClaims{}, fmt.Errorf("Claim 'aud' is not of type string")
	}

	issuedAt, ok := claimsMap["iat"]
	if !ok {
		return models.JwtClaims{}, fmt.Errorf("Claim 'iat' is not of type int64")
	}

	fmt.Println(reflect.TypeOf(expiresAt))

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

package utils

import (
	"Golang/Http/Auth"
	"Golang/Models"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func InternalServerErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func GenerateAuthToken(user *Models.User) (string, error) {
	claims := &Auth.UserClaims{
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "kanban-api",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func ParseToken(tokenString string) (*Auth.UserClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &Auth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Auth.UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func Template(templateName string, data interface{}) (string, error) {
	templatePathSplit := strings.Split(templateName, ".")
	templatePath := templatePathSplit[:len(templatePathSplit)-1]
	confirmationTemplate, err := template.ParseFiles(
		filepath.Join(
			"resources",
			"templates",
			strings.Join(templatePath, "/"),
			templatePathSplit[len(templatePathSplit)-1]+".html",
		),
	)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var body bytes.Buffer
	if err := confirmationTemplate.Execute(&body, data); err != nil {
		log.Println(err)
		return "", err
	}

	return body.String(), nil
}

package Auth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
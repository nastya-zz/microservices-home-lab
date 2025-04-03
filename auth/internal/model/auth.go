package model

import "github.com/dgrijalva/jwt-go"

const (
	ExamplePath = "/user-v1.UserV1/Get"
)

type UserClaims struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Role string `json:"role"`
}

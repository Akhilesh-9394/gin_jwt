package model

import (
	//...
	// import the jwt-go library
	"github.com/dgrijalva/jwt-go"
	//...
)

type User struct {
	Id       int    `json:"id" `
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

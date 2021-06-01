package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fullstacker-go/practice_gin/model"
	"github.com/gin-gonic/gin"
)

var user1 = &model.User{
	Email:    "akhilesh@gmail.com",
	Password: "123456",
}
var jwtKey = []byte("my_secret_key")

func Signin(c *gin.Context) {
	var newuser model.User
	c.Bind(&newuser)
	if newuser.Email != user1.Email || newuser.Password != user1.Password {
		c.JSON(400, gin.H{
			"error": "invalid Email or Password",
		})
		return
	}
	fmt.Println(newuser)
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &model.Claims{
		Username: user1.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(jwtKey)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	c.JSON(http.StatusOK, tokenString)

}

func Signup(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

func HomeHandler(c *gin.Context) {
	token, _ := c.Cookie("token")
	tknStr := token

	// Initialize a new instance of `Claims`
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "hello world",
		"user":    claims.Username,
	})
}

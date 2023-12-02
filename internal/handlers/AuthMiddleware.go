package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("api-key")

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("auth")

	err := TokenCheck(tokenString)
	if err != nil {
		c.JSON(401, err.Error())
	}
	c.Next()
}

func TokenCheck(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("not a valid token")
	}

	return nil
}

package handlers

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type UserData struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	// Add other fields as needed
}

func Authorization(c *gin.Context) {
	var user UserData
	Email := c.GetHeader("email")
	Password := c.GetHeader("password")

	db, err := sql.Open("postgres", os.Getenv("DatabaseDSN"))
	if err != nil {
		log.Fatal("[ERROR] Problem with database connectivity: ", err)
		return
	}

	defer db.Close()

	query := "SELECT id, email, created_at FROM profile WHERE email = $1 AND hash = $2"
	err = db.QueryRow(query, Email, Password).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Wrong Credentials"})
			c.Abort()
			return
		}
		fmt.Println(err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(500, err)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

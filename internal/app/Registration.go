package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"
)

type ProfileType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

func Registration(c *gin.Context) {
	var profile ProfileType
	if err := c.BindJSON(&profile); err != nil {
		c.IndentedJSON(401, "All fields must not be empty")
		return
	}

	if ValidEmail(profile.Email) == false {
		c.IndentedJSON(401, "Not valid email address.")
		return
	}

	if strings.TrimSpace(profile.Password) == "" {
		c.IndentedJSON(400, "Name and Password fields cannot be empty")
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("[ERROR] Problem with env variable: ", err)
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DatabaseDSN"))
	if err != nil {
		log.Fatal("[ERROR] Problem with database connectivity: ", err)
		return
	}

	defer db.Close()

	query := "INSERT INTO profile (email, hash) VALUES($1, $2)"
	if err != nil {
		println("46 line Registration.go: " + err.Error())
	}

	_, err = db.Query(query, profile.Email, profile.Password)
	if err != nil {
		println("52 line Registration.go: " + err.Error())
		c.IndentedJSON(401, "This email was already used")
		return
	}

	c.IndentedJSON(http.StatusOK, "You registered successfully!")
	return
}

package app

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// TODO: add jwt token

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
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(404, gin.H{"error": "Wrong Credentials"})
			return
		}
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, user)

}

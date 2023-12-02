package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func GetProfile(c *gin.Context) {

	token := c.GetHeader("auth")
	Email := c.GetHeader("email")

	err := TokenCheck(token)

	if err != nil {
		c.JSON(401, "Not Authorized")
		c.Abort()
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

	var user User
	query := "SELECT id, email FROM profile WHERE email LIKE '%' || $1 || '%'"
	err = db.QueryRow(query, Email).Scan(&user.ID, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.IndentedJSON(404, "User not found")
			return
		}
		println(err.Error())
		c.IndentedJSON(500, "Database Error")
		return
	}
	c.JSON(200, user)

}

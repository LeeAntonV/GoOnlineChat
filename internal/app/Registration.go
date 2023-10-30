package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strings"
)

type ProfileType struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Registration(c *gin.Context) {
	var profile *ProfileType

	if err := c.BindJSON(&profile); err != nil {
		return
	}

	if strings.TrimSpace(profile.Name) == "" || strings.TrimSpace(profile.Password) == "" {
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

	c.IndentedJSON(http.StatusOK, "You registered successfully!")
	return
}

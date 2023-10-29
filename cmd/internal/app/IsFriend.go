package app

import "github.com/gin-gonic/gin"

func IsFriends(c *gin.Context) {
	id := c.GetHeader("id")

	for _, prof := range ProfileAuth {
		if prof.ID == id {
			if prof.IsFriends == true {
				c.IndentedJSON(200, "You are befriended with "+prof.Name)
				return
			}
			c.IndentedJSON(200, "You are not friends with "+prof.Name)
			return
		}
	}
	c.IndentedJSON(404, "User with id: "+id+"was not found")
}

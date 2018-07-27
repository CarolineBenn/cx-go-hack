package stats

import (
  "github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
  user := c.Params.ByName("id")

  c.JSON(200, gin.H{"user": user, "data": "TODO"})
}
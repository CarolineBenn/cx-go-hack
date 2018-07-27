package endpoints

import (
	"github.com/gin-gonic/gin"
)

// Health func
func Health(c *gin.Context) {
	c.JSON(200, gin.H{"health": "ok"})
}

package web

import (
	"github.com/CarolineBenn/cx-go-hack/app/web/endpoints"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", endpoints.Health)

	return r
}

package web

import (
	"github.com/CarolineBenn/cx-go-hack/app/web/endpoints/main"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/books", main.Books)

	return r
}

package web

import (
	"github.com/CarolineBenn/cx-go-hack/app/web/endpoints"
	"github.com/CarolineBenn/cx-go-hack/app/web/endpoints/books"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health
	r.GET("/health", endpoints.Health)

	// Books
	r.GET("/books/latest", books.Latest)
	r.GET("/books/toBuy", books.ToBuy)

	// Whole series
	r.GET("/series/:id", books.BookRoute) // id is title-with-hyphens
	r.GET("/series", books.All)

	return r
}

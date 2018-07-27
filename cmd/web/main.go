package main

import (
	"github.com/CarolineBenn/cx-go-hack/app/web"
)

func main() {
	router := web.SetupRouter()

	router.Run(":8080")
}

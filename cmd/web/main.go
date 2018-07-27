package main

import (
  "github.com/deliveroo/user-data-service/app/web"
)

func main() {
  router := web.SetupRouter()

  router.Run(":8080")
}
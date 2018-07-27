package web

import (
  "github.com/gin-gonic/gin"
  "github.com/deliveroo/user-data-service/app/web/endpoints/stats"
)

func SetupRouter() *gin.Engine {
  r := gin.Default()

  r.GET("/user/:id", stats.User)

  return r
}

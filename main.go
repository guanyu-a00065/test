package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())

	apis := server.Group("/test/api")
	apis.GET("/v", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Test Successful..")

	})

	server.Run("9512")
}

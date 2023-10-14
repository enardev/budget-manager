package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	app.Run()
}

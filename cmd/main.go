package main

import (
	"signalzero/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	handle := handlers.New()
	
	router.GET("/", handle.GetUsers)
	router.POST("/", handle.AddUsers)

	router.Run(":8080")

}

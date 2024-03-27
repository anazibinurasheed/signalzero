package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	handle := &handler{}

	router.GET("/", handle.GetUsers)
	router.POST("/", handle.AddUsers)

	router.Run(":8080")

}

type handler struct {
}

func (h *handler) GetUsers(c *gin.Context) {
	query := c.Query("username")

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	users, err := FetchUsers(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "unexpected error caused by server",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   users,
	})
}

func (h *handler) AddUsers(c *gin.Context) {
	body := User{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := InsertOne(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "unexpected error caused by server",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"msg":    "success",
	})
}

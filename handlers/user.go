package handlers

import (
	"context"
	"net/http"
	"signalzero/db"
	"signalzero/models"
	"time"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) GetUsers(c *gin.Context) {
	query := c.Query("username")

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	users, err := db.FetchUsers(ctx, query)
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
	body := models.User{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	err := db.InsertOne(ctx, body)
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

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the protected dashboard!",
		"userID":  userID,
		"role":    role,
		"data": []string{
			"Secret data 1",
			"Secret data 2",
		},
	})
}

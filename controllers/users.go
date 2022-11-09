package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "GetUserById - Controller"})
}

func GetUserById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "GetUserById - Controller"})
}

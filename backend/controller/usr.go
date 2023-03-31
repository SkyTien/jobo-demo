package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"userName" binding:"required"`
	Password string `json:"passWord" binding:"required"`
}

func ReqLogin(c *gin.Context) {

	log.Println("TRY LOGIN")
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "abc" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

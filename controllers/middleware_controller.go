package controllers

import (
	"final_project/middleware"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
)


func JWTController(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token invalid or expiry"})
		return
	}
}
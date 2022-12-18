package controllers

import (
	"github.com/gin-gonic/gin"
	"final_project/middleware"
	"strings"
)

func Login(c *gin.Context) {

}

func Show(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	claim, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
	}
	email := claim.Gmail
}
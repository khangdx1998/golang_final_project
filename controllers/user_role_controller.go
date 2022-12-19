package controllers

import (
	"final_project/models"
	"github.com/gin-gonic/gin"
	"strings"
	"final_project/middleware"
	"net/http"
	repo "final_project/repository"
)
func Create_UserRole(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	user_role := models.UserRole{}

	if err := c.BindJSON(&user_role); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	e := repo.CreateUserRole(user_role)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create user_role successfully"})
}

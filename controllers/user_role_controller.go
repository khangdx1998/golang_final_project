package controllers

import (
	"final_project/models"
	"github.com/gin-gonic/gin"
	"net/http"
	repo "final_project/repository"
)
func CreateUserRoleController(c *gin.Context) {
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

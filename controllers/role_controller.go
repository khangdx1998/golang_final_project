package controllers

import (
	repo "final_project/repository"
	"net/http"
	"final_project/models"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func ReadRoleController(c *gin.Context) {
	user, err := repo.ReadRole(models.Condition{Field: "id", Value: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	json_bytes, err := json.Marshal(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	} 
	
	c.Data(http.StatusOK, "application/json", json_bytes)	 
}

func CreateRoleController(c *gin.Context) {
	info := models.Role{}
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	e := repo.CreateRole(info)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create role successfully"})
}

func DeleteRoleController(c *gin.Context) {
	id :=  c.Param("id")

	e := repo.DeleteRole(models.Condition{Field: "id", Value: id})
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully id = " + id})
}

func UpdateRoleController(c *gin.Context) {
	id :=  c.Param("id")

	info := models.Role{}
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	e := repo.UpdateRole(models.Condition{Field: "id", Value: id}, info)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update successfully id = " + id})
}
package controllers

import (
	"final_project/middleware"
	repo "final_project/repository"
	// "fmt"
	"net/http"
	"strings"
	"final_project/models"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func ShowRole(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

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

func CreateRole(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

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

func DeleteRole(c *gin.Context) {
	id :=  c.Param("id")
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	e := repo.DeleteRole(models.Condition{Field: "id", Value: id})
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully id = " + id})
}

func UpdateRole(c *gin.Context) {
	id :=  c.Param("id")
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

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
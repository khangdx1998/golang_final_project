package controllers

import (
	"final_project/middleware"
	repo "final_project/repository"
	"fmt"
	"net/http"
	"strings"
	"final_project/models"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func Login(c *gin.Context) {
	credentials := models.Credentials{}

	if err := c.BindJSON(&credentials); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := repo.ReadUser(models.Condition{Field: "email", Value: credentials.Email})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gmail or password invalid"})
		return
	}
	if credentials.Password != user.Password {
		fmt.Println("Not  match")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Gmail or password invalid"})
		return
	}

	jwt_token, err := middleware.GenerateJWT(credentials.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	jwt_token = "Bearer " + jwt_token
	c.JSON(http.StatusOK, gin.H{"token": jwt_token})
}

func Update(c *gin.Context) {
	id :=  c.Param("id")
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	info := models.User{}
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, e := repo.UpdateUser(models.Condition{Field: "id", Value: id}, info)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update successfully id = " + id})
}

func Delete(c *gin.Context) {
	id :=  c.Param("id")
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	_, e := repo.DeleteUser(models.Condition{Field: "id", Value: id})
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully id = " + id})
}

func Create(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	info := models.User{}
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, e := repo.CreateUser(info)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Create user successfully"})
}

func Show(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	user, err := repo.ReadUser(models.Condition{Field: "id", Value: c.Param("id")})
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

func GetRoles(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	_, err := middleware.DecodeJWT(token)
	if err != nil {
		c.JSON(403, "Cannot decode token")
		return
	}

	roles, err := repo.GetListRoles(models.Condition{Field: "id", Value: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	j, err := json.MarshalIndent(roles, "", " ")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	
	c.Data(http.StatusOK, "application/json", j)
}
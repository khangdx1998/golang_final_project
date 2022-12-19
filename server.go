package main

import (
	"final_project/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	group1 := r.Group("/api")
	{	
		// Login
		group1.GET("/login", controllers.Login)

		client := group1.Group("/users")
		{
			client.PUT("/update/:id", controllers.Update)
			client.DELETE("/delete/:id", controllers.Delete)
			client.POST("/create", controllers.Create)
		}
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}






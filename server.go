package main

import (
	"final_project/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	group1 := r.Group("/api")
	{
		client := group1.Group("/users")
		{
			client.GET("/", controllers.Show)
		}
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}






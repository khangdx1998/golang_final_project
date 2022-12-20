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

		users := group1.Group("/users")
		{
			users.PUT("/update/:id", controllers.Update)
			users.DELETE("/delete/:id", controllers.Delete)
			users.POST("/create", controllers.Create)
			users.GET("/get/:id", controllers.Show)
			users.GET("/get_roles/:id", controllers.GetRoles)
		}

		role := group1.Group("/role")
		{
			role.GET("/get/:id", controllers.ShowRole)
			role.POST("/create", controllers.CreateRole)
			role.DELETE("/delete/:id", controllers.DeleteRole)
			role.PUT("/update/:id", controllers.UpdateRole)
		}

		user_role := group1.Group("/user_role")
		{
			user_role.POST("/create", controllers.Create_UserRole)
		}
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}






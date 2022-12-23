package router

import (
	"final_project/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	group1 := r.Group("/api")
	{	
		// Login
		group1.GET("/login", controllers.Login)
		// User
		users := group1.Group("/users").Use(controllers.JWTController)
		{
			users.PUT("/update/:id", controllers.UpdateUserController)
			users.DELETE("/delete/:id", controllers.DeleteUserController)
			users.POST("/create", controllers.CreateUserController)
			users.GET("/get/:id", controllers.ReadUserController)
			users.GET("/get_roles/:id", controllers.GetRolesController)
		}
		// Role
		role := group1.Group("/role").Use(controllers.JWTController)
		{
			role.GET("/get/:id", controllers.ReadRoleController)
			role.POST("/create", controllers.CreateRoleController)
			role.DELETE("/delete/:id", controllers.DeleteRoleController)
			role.PUT("/update/:id", controllers.UpdateRoleController)
		}
		//User Role
		user_role := group1.Group("/user_role").Use(controllers.JWTController)
		{
			user_role.POST("/create", controllers.CreateUserRoleController)
		}
	}
	return r
}
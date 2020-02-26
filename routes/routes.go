package routes

import (
	"github.com/gin-gonic/gin"
	user "github.com/mfirmanakbar/gin-beego-mux-buffalo/controllers/users"
	"net/http"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}

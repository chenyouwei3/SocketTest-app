package router

import (
	"SocketTest-app/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("/login", controller.LoginUser)
		user.POST("/create", controller.RegisterUser)
		user.DELETE("/delete", controller.DeleteUser)
		user.PUT("/update", controller.UpdateUser)
		user.GET("/get", controller.GetUserData)
	}
}

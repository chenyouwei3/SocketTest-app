package controller

import (
	"SocketTest-app/model"
	"SocketTest-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.LoginUser(user))
}

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.RegisterUser(user))
}
func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.DeletedUser(id))
}
func UpdateUser(c *gin.Context) {
	flag := c.DefaultQuery("flag", "false")
	if flag != "false" && flag != "true" {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "参数错误")
		return
	}
	c.JSON(http.StatusOK, service.UpdateUser(user, flag))
}
func GetUser(c *gin.Context) {
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "0")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetUser(currPage, pageSize, startTime, endTime))
}

///dadad

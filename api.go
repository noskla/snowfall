package main

import (
	"github.com/gin-gonic/gin"
)

func InitAPIRouter(router *gin.Engine) {
	APIRouter := router.Group("/api")
	{
		APIRouter.POST("/user", routeCreateUser)
	}
}

type routeCreateUserPOSTBody struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Discord  string `form:"discord" json:"discord" binding:"required"`
}
func routeCreateUser(c *gin.Context) {
	var data routeCreateUserPOSTBody
	err := c.ShouldBind(&data)
	if errorOccurred(err, false) {
		c.JSON(400, gin.H{
			"success": false,
			"reason": "Missing arguments",
		})
		return
	}

	ok, res := createUser(data.Username, data.Password, data.Discord)
	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"reason": res,
		})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"uuid": res,
		})
	}
}

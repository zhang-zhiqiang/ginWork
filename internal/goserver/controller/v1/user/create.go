package user

import (
	v1 "baseframe/api/goserver/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *UserController) Create(c *gin.Context) {

	var req *v1.CreateUserReq

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bind error",
		})
		return
	}

	if err := uc.us.CreateUser(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

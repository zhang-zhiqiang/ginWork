package user

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/internal/pkg/code"
	"baseframe/pkg/core"
	"github.com/gin-gonic/gin"
)

func (uc *UserController) Create(c *gin.Context) {

	var req *v1.CreateUserReq

	if err := c.ShouldBindJSON(req); err != nil {
		core.WriteResponse(c, code.ErrBind, nil)
		return
	}

	if err := uc.us.CreateUser(c, req); err != nil {
		core.WriteResponse(c, code.ErrSuccess, nil)
		return
	}

	core.WriteResponse(c, code.ErrSuccess, nil)
}

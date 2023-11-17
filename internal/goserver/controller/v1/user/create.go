package user

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/internal/pkg/code"
	"baseframe/pkg/core"
	"baseframe/pkg/errors"
	"baseframe/pkg/validation"
	"github.com/gin-gonic/gin"
)

func (uc *UserController) Create(c *gin.Context) {

	var req v1.CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		core.Response(c, errors.WithCode(code.ErrBind, err.Error()), nil)
		return
	}

	if err := validation.Check(&req); err != nil {
		core.Response(c, errors.WithCode(code.ErrValidation, err.Error()), nil)
		return
	}

	if err := uc.us.CreateUser(c, &req); err != nil {
		core.Response(c, errors.WithCode(code.ErrSuccess, err.Error()), nil)
		return
	}

	core.Response(c, nil, nil)
}

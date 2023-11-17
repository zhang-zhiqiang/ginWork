package core

import (
	"baseframe/pkg/errors"
	"baseframe/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var codeSuccess = 100000

func success(data interface{}) ErrResponse {
	return ErrResponse{
		Code:    codeSuccess,
		Message: "ok",
		Data:    data,
	}
}

func Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		log.Errorf("错误信息： %#+v", err)
		coder := errors.ParseCoder(err)

		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:    coder.Code(),
			Message: coder.String(),
			Data:    data,
		})

		return
	}

	c.JSON(http.StatusOK, success(data))
}

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, code int, data interface{}) {
	if code != 0 {
		coder := GetCoder(code)

		c.JSON(coder.HttpStatus(), ErrResponse{
			Code:    coder.Code(),
			Message: coder.Message(),
			Data:    data,
		})

		return
	}

	c.JSON(http.StatusOK, ErrResponse{Data: data})
}

func WithMsgResponse(c *gin.Context, code int, message string, data interface{}) {
	if code != 0 {
		coder := GetCoder(code)

		c.JSON(coder.HttpStatus(), ErrResponse{
			Code:    coder.Code(),
			Message: message,
			Data:    data,
		})

		return
	}

	c.JSON(http.StatusOK, ErrResponse{Data: data})
}

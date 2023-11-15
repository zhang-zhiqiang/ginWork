package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Reference string      `json:"reference,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func WriteResponse(c *gin.Context, code int, data interface{}) {
	if code != 0 {
		coder := GetCoder(code)

		c.JSON(coder.HttpStatus(), Response{
			Code:    coder.Code(),
			Message: coder.Message(),
			Data:    data,
		})
	}

	c.JSON(http.StatusOK, Response{Data: data})
}

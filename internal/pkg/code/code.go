package code

import (
	"baseframe/pkg/core"
	"baseframe/pkg/log"
	"github.com/novalagung/gubrak"
)

func register(code int, httpStatus int, msg string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		httpStatus = 400
		log.Errorf("错误码注册 状态码越界 code %d httpStatus %d msg %s", code, httpStatus, msg)
	}
	coder := core.DefaultCoder{C: code, Hs: httpStatus, Msg: msg}
	core.MustRegister(coder)
}

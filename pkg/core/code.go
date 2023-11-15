package core

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	unknowCoder DefaultCoder = DefaultCoder{1, http.StatusInternalServerError, "异常错误码"}
)

func init() {
	codes[unknowCoder.C] = unknowCoder
}

type Coder interface {
	Code() int
	HttpStatus() int
	Message() string
}

type DefaultCoder struct {
	C   int
	Hs  int
	Msg string
}

func (coder DefaultCoder) Code() int {
	return coder.C
}

func (coder DefaultCoder) HttpStatus() int {
	return coder.Hs
}

func (coder DefaultCoder) Message() string {
	return coder.Msg
}

var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic(fmt.Sprintf("错误码不能为0 %s", coder.Message()))
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("错误码已经存在 %d %s", coder.Code(), coder.Message()))
	}

	codes[coder.Code()] = coder
}

func GetCoder(code int) Coder {
	if coder, ok := codes[code]; ok {
		return coder
	}
	return unknowCoder
}

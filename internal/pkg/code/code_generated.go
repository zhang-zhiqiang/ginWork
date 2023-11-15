// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by "codegen -type=int"; DO NOT EDIT.

package code

// init register error codes defines in this source code to `github.com/marmotedu/errors`
func init() {
	register(ErrSuccess, 200, "OK")
	register(ErrUnknown, 500, "Internal server error")
	register(ErrBind, 400, "绑定参数异常")
	register(ErrValidation, 400, "校验失败")
	register(ErrTokenInvalid, 401, "令牌无效")
	register(ErrPageNotFound, 404, "Page not found")
}
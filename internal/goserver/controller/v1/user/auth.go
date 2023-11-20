package user

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/internal/goserver/conf"
	"baseframe/internal/pkg/code"
	"baseframe/pkg/core"
	"baseframe/pkg/log"
	"baseframe/pkg/validation"
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewJwtAuth(config *conf.Config, uc *UserController) *jwt.GinJWTMiddleware {
	ginJwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.Jwt.Realm,
		SigningAlgorithm: "HS256",
		Key:              []byte(config.Jwt.Key),
		Timeout:          config.Jwt.Timeout,
		MaxRefresh:       config.Jwt.MaxRefresh,
		Authenticator:    uc.authenticator(),
		LoginResponse:    uc.loginResponse(),
		Unauthorized:     uc.unauthorized(),
		RefreshResponse:  uc.refreshResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
	})

	return ginJwt
}

var (
	loginErrFailed  = errors.New("用户名或密码错误")
	loginErrMissing = errors.New("用户名或密码为空")
)

func (uc *UserController) authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var req v1.LoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Errorf("登陆解析: %s", err.Error())

			return nil, loginErrMissing
		}

		if err := validation.Check(&req); err != nil {
			return nil, loginErrFailed
		}

		user, err := uc.us.Login(c, &req)
		if err != nil {
			return nil, loginErrFailed
		}

		return user, nil
	}
}

func (uc *UserController) unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, status int, message string) {
		if status == http.StatusOK {
			core.Response(c, nil, nil)
		} else {
			failed := core.ErrResponse{
				Code:    code.ErrLoginFailed,
				Message: message,
			}
			c.JSON(status, failed)
		}
	}
}

func (uc *UserController) loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		response := &v1.LoginRep{
			Token:  token,
			Expire: expire.Format(time.RFC3339),
		}
		core.Response(c, nil, response)
	}
}

func (uc *UserController) refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		response := &v1.LoginRep{
			Token:  token,
			Expire: expire.Format(time.RFC3339),
		}
		core.Response(c, nil, response)
	}
}

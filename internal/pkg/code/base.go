package code

const (
	// ErrSuccess - 200: OK.
	ErrSuccess int = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrBind - 400: 绑定参数异常.
	ErrBind

	// ErrValidation - 400: 校验失败.
	ErrValidation

	// ErrTokenInvalid - 401: 令牌无效.
	ErrTokenInvalid

	// ErrPageNotFound - 404: Page not found.
	ErrPageNotFound
)

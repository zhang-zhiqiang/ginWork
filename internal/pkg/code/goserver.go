package code

const (
	// ErrUserNotFound - 404: 用户不存在.
	ErrUserNotFound int = iota + 110001

	// ErrLoginFailed - 401: 登陆失败.
	ErrLoginFailed
)

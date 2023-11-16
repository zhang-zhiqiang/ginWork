package v1

type CreateUserReq struct {
	Username string `validate:"required,min=1,max=128" label:"用户名"`
	Nickname string `validate:"required,min=1,max=128" label:"昵称"`
	Password string `validate:"required,min=1,max=16" label:"密码"`
	Email    string `validate:"required,email,min=1,max=100" label:"邮箱"`
	Phone    string `validate:"required,phone" label:"手机号"`
	Status   int    `validate:"omitempty" label:"状态"`
}

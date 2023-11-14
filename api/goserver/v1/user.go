package v1

type CreateUserReq struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
}

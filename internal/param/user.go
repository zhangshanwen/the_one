package param

type (
	Register struct {
		Mobile   string `json:"mobile"     binding:"required"`
		NickName string `json:"nick_name"  binding:"required"`
		Password string `json:"password"   binding:"required"`
		Email    string `json:"email"      binding:"required"`
		Avatar   string `json:"avatar"`
	}
	Login struct {
		Mobile   string `json:"mobile"     binding:"required"`
		Password string `json:"password"   binding:"required"`
	}
)

package response

type (
	UserInfo struct {
		Authorization string `json:"authorization,omitempty"`
		Mobile        string `json:"mobile"`
		NickName      string `json:"nick_name"`
		Email         string `json:"email"`
		Avatar        string `json:"avatar"`
		Balance       int64  `json:"balance"`
	}
)

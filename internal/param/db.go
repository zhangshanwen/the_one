package param

type CreateDatabase struct {
	Database string `json:"database"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

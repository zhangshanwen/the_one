package conf

type DB struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}
type Mysql struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	DataBase    string `yaml:"dataBase"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	LogMode     string `yaml:"log_mode"`
	Config      string `yaml:"config"`
	TablePrefix string `yaml:"table_prefix"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DataBase int    `yaml:"dataBase"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
}

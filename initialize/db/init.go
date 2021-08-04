package db

func InitDb() {
	InitMysql() // s初始化mysql
	InitRedis() // 初始化redis
}

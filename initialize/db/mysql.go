package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/zhangshanwen/the_one/initialize/conf"
	l "github.com/zhangshanwen/the_one/initialize/logger"
	"github.com/zhangshanwen/the_one/model"
)

var G *gorm.DB

func InitMysql() {
	l.Logger.Info("--------init_mysql_client_end---------")
	var err error
	m := conf.C.DB.Mysql
	newLogger := logger.New(
		log.New(l.Logger.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DataBase, m.Config)
	if G, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.TablePrefix,
			SingularTable: true,
		},
	}); err != nil {
		panic(err)
	}
	l.Logger.Info("--------init_mysql_client_end---------")
	AutoMigrate()
	return
}
func AutoMigrate() {
	l.Logger.Info("--------mysql_auto_migrate_start---------")
	_ = G.AutoMigrate(model.User{}, model.Wallet{})
	l.Logger.Info("--------mysql_auto_migrate_end---------")
}

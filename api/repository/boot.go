package repository

import (
	"database/sql"
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbClient *gorm.DB

func Boot() {
	dbConnect, err := sql.Open(viper.GetString("db.driver-name"), viper.GetString("db.url"))
	dbConnect.SetMaxIdleConns(5)
	dbConnect.SetMaxOpenConns(10)
	exception.Check(err)
	dbClient, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbConnect,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

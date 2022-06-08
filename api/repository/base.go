package repository

import (
	"database/sql"
	"fmt"
	"github.com/czyhome/go-pulsar-manager/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbClient *gorm.DB

func init() {
	var err error
	dbConnect, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/go-pulsar-manager?charset=utf8mb4&parseTime=True", "admin", "***REMOVED***", "192.168.2.18", "3306"))
	dbConnect.SetMaxIdleConns(5)
	dbConnect.SetMaxOpenConns(10)
	web.CheckError(err)
	dbClient, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbConnect,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

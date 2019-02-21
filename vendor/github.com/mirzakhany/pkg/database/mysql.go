package database

import (
	"fmt"

	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mirzakhany/pkg/logger"
)

// InitMySQLDB init mySQL
func InitMySQLDB(maxTry int, settings MySQLSettings) (*gorm.DB, error) {

	var err error
	var database *gorm.DB

	net := "tcp"
	address := fmt.Sprintf("%s:%d", settings.Host, settings.Port)
	if settings.Socket != "" {
		net = "unix"
		address = settings.Socket
	}

	configs := mysql.NewConfig()
	configs.Params = settings.Options
	configs.User = settings.Username
	configs.Passwd = settings.Password
	configs.DBName = settings.DatabaseName
	configs.Net = net
	configs.Addr = address
	configs.WriteTimeout = time.Duration(settings.WriteTimeout) * time.Microsecond
	configs.ReadTimeout = time.Duration(settings.ReadTimeout) * time.Microsecond
	configs.Timeout = time.Duration(settings.Timeout) * time.Microsecond

	for {
		database, err = gorm.Open("mysql", configs.FormatDSN())
		if err == nil {
			break
		}
		logger.Errorf("Connect to MySql failed du error: %v - retrying ...", err)
		if maxTry > 0 {
			maxTry--
		} else {
			return nil, err
		}
	}
	logger.Info("Connection to MySql stabilised")
	return database, err
}

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mirzakhany/pkg/database/postgresql"
	"github.com/mirzakhany/pkg/logger"
	"time"
)

// InitPostgres init postgres
func InitPostgres(maxTry int, settings PostgresSettings) (*gorm.DB, error) {

	var err error
	var database *gorm.DB

	configs := postgresql.ConnectionURL{
		User:     settings.Username,
		Password: settings.Password,
		Host:     settings.Host,
		Socket:   settings.Socket,
		Database: settings.DatabaseName,
		Options:  settings.Options,
	}

	for {
		database, err = gorm.Open("postgres", configs.String())
		if err == nil {
			break
		}
		logger.Errorf("Connect to Postgres failed du error: %v - retrying ...", err)
		if maxTry > 0 {
			maxTry--
			time.Sleep(time.Second * 1)
		} else {
			return nil, err
		}
	}
	logger.Info("Connection to Postgres stabilised")
	return database, err
}

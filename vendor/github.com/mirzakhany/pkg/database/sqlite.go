package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mirzakhany/pkg/logger"
)

// InitSQLite init SQLite
func InitSQLite(maxTry int, settings SQLiteSettings) (*gorm.DB, error) {

	var err error
	var db *gorm.DB
	for {
		db, err = gorm.Open("sqlite3", settings.DatabaseName)
		if err == nil {
			break
		}
		logger.Errorf("Connect to SQLite failed du error: %v - retrying ...", err)
		if maxTry > 0 {
			maxTry--
		} else {
			return nil, err
		}
	}
	logger.Info("Connection to SQLite stabilised")
	return db, err
}

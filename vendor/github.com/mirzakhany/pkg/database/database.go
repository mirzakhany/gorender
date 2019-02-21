package database

import (
	"github.com/jinzhu/gorm"
	"github.com/mirzakhany/pkg/logger"
	"github.com/pkg/errors"
	"gopkg.in/gormigrate.v1"
)

var (
	// DB database instance
	DB *gorm.DB
)

// KV key value type
type KV map[string]string

// DBSettings settings of database
type DBSettings struct {
	Engine       string           `json:"engine"`
	MaxTry       int              `json:"max_try"`
	MySQL        MySQLSettings    `json:"mysql"`
	Postgres     PostgresSettings `json:"postgres"`
	SQLite       SQLiteSettings   `json:"sqlite"`
	DBModels     []interface{}
	DBMigrations []*gormigrate.Migration
	ForeignKeys  []ForeignKey
}

// MySQLSettings is settings of mySQL
type MySQLSettings struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Socket       string `json:"socket"`
	DialTimeout  int    `json:"dial_timeout"`
	Options      KV     `json:"options"`
	Timeout      int    `json:"timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

// PostgresSettings is settings of Postgres
type PostgresSettings struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Socket       string `json:"socket"`
	DialTimeout  int    `json:"dial_timeout"`
	Options      KV     `json:"options"`
}

// SQLiteSettings is settings of SQLite
type SQLiteSettings struct {
	DatabaseName string `json:"database_name"`
}

// ForeignKey struct of table foreign keys
type ForeignKey struct {
	Model       interface{}
	Field       string
	Destination string
	OnDelete    string
	OnUpdate    string
}

// InitDatabase for initialize database
func InitDatabase(settings *DBSettings) (*gorm.DB, error) {
	logger.LogAccess.Infof("Init Database Engine as %s", settings.Engine)
	var err error
	switch settings.Engine {
	case "postgres":
		DB, err = InitPostgres(settings.MaxTry, settings.Postgres)
	case "sqlite":
		DB, err = InitSQLite(settings.MaxTry, settings.SQLite)
	case "mysql":
		DB, err = InitMySQLDB(settings.MaxTry, settings.MySQL)
	default:
		logger.Error("database error: can't find database driver")
		return nil, errors.New("can't find database driver")
	}

	err = initSchema(DB, settings)
	return DB, err
}

func initSchema(db *gorm.DB, settings *DBSettings) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, settings.DBMigrations)
	m.InitSchema(func(tx *gorm.DB) error {
		if len(settings.DBModels) > 0 {
			err := tx.AutoMigrate(
				settings.DBModels...,
			).Error
			if err != nil {
				logger.Fatalf("auto migration failed, %s", err)
				return errors.New("auto migration failed")
			}
		}
		for _, fk := range settings.ForeignKeys {
			if err := tx.Model(fk.Model).AddForeignKey(fk.Field, fk.Destination, fk.OnDelete, fk.OnUpdate).Error; err != nil {
				logger.Fatalf("add foreign-key failed, %v", err)
				return err
			}
		}
		return nil
	})
	err := m.Migrate()
	return err
}

// CloseDatabase close database session
func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		logger.Errorf("database error: can't close database, %v", err)
	}
}

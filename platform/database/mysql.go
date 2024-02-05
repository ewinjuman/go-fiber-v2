package database

import (
	"fmt"
	"go-fiber-v2/pkg/configs"
	Session "go-fiber-v2/pkg/libs/session"
	"go-fiber-v2/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

//func init() {
//	err := mysqlOpen()
//	if err != nil {
//		panic(err.Error())
//	}
//}

// Mysql open connection
func mysqlOpen() error {
	//var err error
	config := configs.Config.Database

	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return err
	}

	db, err := gorm.Open(mysql.Open(mysqlConnURL))
	if err != nil {
		//fmt.Println("Failed to connect database. err: ", err.Error())
		return fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)

	DB = db
	return nil
}

// MysqlConnection func for connection to Mysql database.
func MysqlConnection(session *Session.Session) (*gorm.DB, error) {
	if DB == nil {
		if err := mysqlOpen(); err != nil {
			session.Error(err.Error())
			return DB, err
		}
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	if errping := sqlDB.Ping(); errping != nil {
		errping = nil
		if errping = mysqlOpen(); errping != nil {
			session.Error(errping.Error())
			return DB, errping
		}
	}
	logLevel := logger.Info
	if !configs.Config.Database.LogMode {
		logLevel = logger.Silent
	}
	newLogger := logger.New(
		session.Logger, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	//DB.Logger.LogMode(logger.Silent)
	DB.Logger = newLogger
	return DB, nil
}

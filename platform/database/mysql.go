package database

import (
	"fmt"
	Session "github.com/ewinjuman/go-lib/session"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/repository"
	"go-fiber-v2/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	dbConnection *gorm.DB
	once         sync.Once
)

//func init() {
//	err := mysqlOpen()
//	if err != nil {
//		panic(err.Error())
//	}
//}

// InitLogger inisialisasi logger sekali saja
func InitDb() error {
	once.Do(func() {
		err := mysqlOpen
		if err != nil {
			panic(err)
		}
	})
	return nil
}

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

	dbConnection = db
	return nil
}

// GetMysqlConnection func for connection to Mysql database.
func GetMysqlConnection(session *Session.Session) (*gorm.DB, error) {
	if dbConnection == nil {
		if err := mysqlOpen(); err != nil {
			session.Error(err.Error())
			return nil, repository.UndefinedErr
		}
	}

	//if err := checkAndPingDatabase(session); err != nil {
	//	return nil, err
	//}

	configureDatabaseLogger(session)

	return dbConnection, nil
}
func checkAndPingDatabase(session *Session.Session) error {
	sqlDB, err := dbConnection.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		if err := mysqlOpen(); err != nil {
			session.Error(err.Error())
			return repository.UndefinedErr
		}
	}
	return nil
}

func configureDatabaseLogger(session *Session.Session) {
	logLevel := logger.Info
	if !configs.Config.Database.LogMode {
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		session.Logger,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)

	dbConnection.Logger = newLogger
}

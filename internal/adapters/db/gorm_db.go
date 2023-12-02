package db

import (
	"e-student/internal/app"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(cfg *app.Config) (*gorm.DB, func() error) {
	logLevel := logger.Silent

	if cfg.SqlLogInfo {
		logLevel = logger.Info
	}

	queryLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	conn, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{Logger: queryLogger})

	if err != nil {
		panic(err)
	}

	sqlDB, err := conn.DB()

	if err != nil {
		panic(err)
	}

	return conn, sqlDB.Close
}

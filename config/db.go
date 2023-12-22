package config

import (
	"log"
	"os"
	"stranger-words/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var LOGGER = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,          // Don't include params in the SQL log
		Colorful:                  false,         // Disable color
	},
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: LOGGER,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Author{}, &models.Word{})
}
func GetDB() *gorm.DB {
	return db
}

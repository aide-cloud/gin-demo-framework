package data

import (
	"gin-demo-framework/config"
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitDB(c *config.DBConfig) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(c.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		slog.Error("DB connect failed", "error", err)
		return
	}
	_db = db.Debug()
}

func GetDB() *gorm.DB {
	return _db
}

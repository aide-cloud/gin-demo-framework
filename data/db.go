package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("/Users/hubiao/me/gin-demo-framework/gin_demo.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	_db = db
}

func GetDB() *gorm.DB {
	return _db
}

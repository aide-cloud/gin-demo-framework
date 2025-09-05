package main

import (
	"gin-demo-framework/data"
	"gin-demo-framework/data/model"
	"log/slog"
	"os"
)

func main() {
	slog.Info("Migrate database")
	db := data.GetDB()
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		slog.Error("Migrate database failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Migrate database success")
}

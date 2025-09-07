package main

import (
	"gin-demo-framework/config"
	"gin-demo-framework/data"
	"gin-demo-framework/data/model"
	"gin-demo-framework/route"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var addr string
var configPath string

var rootCmd = &cobra.Command{
	Use:   "gin-server",
	Short: "This is a demo for gin framework",
	Long:  "This is a demo for gin framework",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig(configPath)
		if err != nil {
			slog.Error("Load config failed", "error", err)
			os.Exit(1)
		}

		data.Init(c)
		slog.Info("Start gin server")
		r := route.New(&c.Server)
		r.Run(addr)
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  "Migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Migrate database")
		db := data.GetDB()
		if err := db.AutoMigrate(&model.Task{}, &model.Log{}); err != nil {
			slog.Error("Migrate database failed", "error", err)
			os.Exit(1)
		}
		slog.Info("Migrate database success")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().StringVarP(&addr, "addr", "a", ":8080", "server address")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "config file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Execute rootCmd failed", "error", err)
		os.Exit(1)
	}
}

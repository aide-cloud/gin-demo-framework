package data

import "gin-demo-framework/config"

func Init(c *config.Config) {
	InitDB(&c.DB)
	InitRDB(&c.Redis)
}

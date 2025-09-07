package model

import "time"

type Log struct {
	ID       int           `json:"id" gorm:"primaryKey;type:bigint(20);autoIncrement"`
	Status   int           `json:"status" gorm:"type:int(11);comment:状态"`
	Latency  time.Duration `json:"latency" gorm:"type:bigint(20);comment:延迟"`
	Path     string        `json:"path" gorm:"type:varchar(255);comment:路径"`
	Method   string        `json:"method" gorm:"type:varchar(255);comment:方法"`
	ClientIP string        `json:"clientIP" gorm:"type:varchar(255);comment:客户端IP"`
}

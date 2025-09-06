package model

import (
	"gin-demo-framework/pkg/cnst"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        int            `json:"id" gorm:"primaryKey;type:bigint(20);autoIncrement"`
	Title     string         `json:"title" gorm:"not null;type:varchar(64);comment:任务标题;uniqueIndex"`
	Remark    string         `json:"remark" gorm:"type:varchar(255);comment:任务备注;default:'这个人很懒，什么都没有留下'"`
	Completed cnst.Status    `gorm:"default:0;comment:任务完成状态;type:tinyint(1)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Task) TableName() string {
	return "todos"
}

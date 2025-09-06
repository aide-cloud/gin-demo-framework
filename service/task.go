package service

import (
	"context"
	"errors"
	"gin-demo-framework/data/model"
	"log/slog"

	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

type TaskListParams struct {
	Keyword string `form:"keyword"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
}

func (t *TaskService) List(ctx context.Context, params *TaskListParams) ([]*model.Task, int64, error) {
	var tasks []*model.Task
	offset := (params.Page - 1) * params.Limit
	db := t.db.Limit(params.Limit).Offset(offset)
	if params.Keyword != "" {
		db = db.Where("title like ?", "%"+params.Keyword+"%")
	}
	// db = db.Table(model.Task{}.TableName()+"_" + hash(params.OrderNo))
	db = db.Model(&model.Task{})
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Find(&tasks).Error; err != nil {
		slog.Error("List tasks failed", "error", err)
		return nil, 0, err
	}

	return tasks, total, nil
}

func (t *TaskService) Detail(ctx context.Context, id int) (*model.Task, error) {
	var task model.Task
	if err := t.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("查询的任务不存在, 请检查输入的ID是否正确")
		}
		slog.Error("查询任务详情失败", "error", err)
		return nil, errors.New("系统错误， 请联系管理员")
	}

	return &task, nil
}

func (t *TaskService) Create(ctx context.Context, task *model.Task) error {

	var err error

	return err
}

func (t *TaskService) UpdateTask(ctx context.Context, task *model.Task, id int) error {
	var err error

	return err
}

func (t *TaskService) DeleteTask(ctx context.Context, id int) error {
	var err error

	return err
}

package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-demo-framework/data/model"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TaskService struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewTaskService(db *gorm.DB, rdb *redis.Client) *TaskService {
	return &TaskService{
		db:  db,
		rdb: rdb,
	}
}

type TaskListParams struct {
	Keyword string `form:"keyword" uri:"keyword" json:"keyword"`
	Page    int    `form:"page" uri:"page" json:"page"`
	Limit   int    `form:"limit" uri:"limit" json:"limit"`
}

type TaskListReply struct {
	Tasks []*model.Task `json:"tasks"`
	Total int64         `json:"total"`
}

func (t *TaskService) List(ctx context.Context, params *TaskListParams) (*TaskListReply, error) {
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
		return nil, err
	}
	if err := db.Find(&tasks).Error; err != nil {
		slog.Error("List tasks failed", "error", err)
		return nil, err
	}

	return &TaskListReply{
		Tasks: tasks,
		Total: total,
	}, nil
}

func (t *TaskService) Detail(ctx context.Context, id int) (*model.Task, error) {
	key := fmt.Sprintf("task:%d", id)
	// 判断redis中是否存在
	count, err := t.rdb.Exists(ctx, key).Result()
	if err != nil {
		slog.Error("查询任务详情失败", "error", err)
		return nil, errors.New("系统错误， 请联系管理员")
	}
	if count > 0 {
		slog.Info("从redis中查询任务详情", "id", id)
		taskBytes, err := t.rdb.Get(ctx, key).Bytes()
		if err != nil {
			slog.Error("查询任务详情失败", "error", err)
			return nil, errors.New("系统错误， 请联系管理员")
		}
		var task model.Task
		json.Unmarshal(taskBytes, &task)
		return &task, nil
	}

	slog.Info("从数据库中查询任务详情", "id", id)
	var task model.Task
	if err := t.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("查询的任务不存在, 请检查输入的ID是否正确")
		}
		slog.Error("查询任务详情失败", "error", err)
		return nil, errors.New("系统错误， 请联系管理员")
	}
	taskBytes, _ := json.Marshal(task)
	t.rdb.Set(ctx, key, taskBytes, 1*time.Minute)
	slog.Info("将任务详情缓存到redis中", "id", id)
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

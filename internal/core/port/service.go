package port

import (
	"context"
	"kn-assignment/internal/core/domain"
)

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) error
	GetTasksByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error)
	UpdateTaskStatus(ctx context.Context, taskID, status string) error
	GetAllTasks(ctx context.Context, filter map[string]string, sort string) ([]domain.Task, error)
	GetTaskSummary(ctx context.Context) ([]domain.TaskSummary, error)
}

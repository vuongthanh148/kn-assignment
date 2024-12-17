package tasksvc

import (
	"context"
	"kn-assignment/internal/core/domain"
)

func (s *service) CreateTask(ctx context.Context, task domain.Task) error {
	return s.repo.CreateTask(ctx, task)
}

func (s *service) GetTasksByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	return s.repo.GetTasksByAssignee(ctx, assigneeID)
}

func (s *service) UpdateTaskStatus(ctx context.Context, taskID, status string) error {
	return s.repo.UpdateTaskStatus(ctx, taskID, status)
}

func (s *service) GetAllTasks(ctx context.Context, filter map[string]string, sort string) ([]domain.Task, error) {
	return s.repo.GetAllTasks(ctx, filter, sort)
}

func (s *service) GetTaskSummary(ctx context.Context) ([]domain.TaskSummary, error) {
	return s.repo.GetTaskSummary(ctx)
}

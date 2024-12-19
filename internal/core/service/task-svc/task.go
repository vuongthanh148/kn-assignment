package tasksvc

import (
	"context"
	"kn-assignment/internal/constant"
	"kn-assignment/internal/core/domain"
	errors "kn-assignment/internal/core/error"
)

func (s *service) CreateTask(ctx context.Context, task domain.CreateTaskRequest, userId string) error {
	if task.Title == "" {
		return errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Title is required")
	}
	return s.taskRepo.CreateTask(ctx, task, userId)
}

func (s *service) AssignTask(ctx context.Context, taskID, assigneeID string) error {
	if taskID == "" || assigneeID == "" {
		return errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Task ID and Assignee ID are required")
	}
	// Validate that the assignee exists and is an employee
	assignee, err := s.userRepo.GetUserByID(ctx, assigneeID)
	if err != nil {
		return errors.NewCustomErrorWithMessage(constant.ErrCodeNotFound, "Assignee not found")
	}
	if assignee.Role != domain.RoleEmployee {
		return errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Assignee must be an employee")
	}
	return s.taskRepo.AssignTask(ctx, taskID, assigneeID)
}

func (s *service) GetTasksByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	if assigneeID == "" {
		return nil, errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Assignee ID is required")
	}
	return s.taskRepo.GetTasksByAssignee(ctx, assigneeID)
}

func (s *service) UpdateTaskStatus(ctx context.Context, taskID string, status domain.TaskStatus, userId string) error {
	if taskID == "" || status == "" {
		return errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Task ID and status are required")
	}
	return s.taskRepo.UpdateTaskStatus(ctx, taskID, status, userId)
}

func (s *service) GetAllTasks(ctx context.Context, userRole, userID string, filter map[string]string, sort string) ([]domain.Task, error) {
	if userRole == string(domain.RoleEmployee) {
		filter["assignee_id"] = userID
	}
	return s.taskRepo.GetAllTasks(ctx, filter, sort)
}

func (s *service) GetTaskSummary(ctx context.Context) ([]domain.TaskSummary, error) {
	return s.taskRepo.GetTaskSummary(ctx)
}

func (s *service) VerifyTaskAssignment(ctx context.Context, taskID, userID string) (bool, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, taskID)
	if err != nil {
		return false, err
	}
	if task.AssigneeID == nil {
		return false, nil
	}
	return *task.AssigneeID == userID, nil
}

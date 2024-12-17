package tasksvc

import "kn-assignment/internal/core/port"

type service struct {
	repo port.TaskRepository
}

func New(taskRepository port.TaskRepository) port.TaskService {
	return &service{repo: taskRepository}
}

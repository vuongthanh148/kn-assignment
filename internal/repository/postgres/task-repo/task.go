package taskrepo

import (
	"context"
	"fmt"
	"kn-assignment/internal/core/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) CreateTask(ctx context.Context, task domain.Task) error {
	query := `INSERT INTO tasks (id, title, description, assignee_id, status, created_at, due_date) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.dbPool.Exec(ctx, query, task.ID, task.Title, task.Description, task.AssigneeID, task.Status, task.CreatedAt, task.DueDate)
	return err
}

func (r *repository) GetTasksByAssignee(ctx context.Context, assigneeID string) ([]domain.Task, error) {
	query := `SELECT * FROM tasks WHERE assignee_id = $1`
	var tasks []domain.Task
	err := pgxscan.Select(ctx, r.dbPool, &tasks, query, assigneeID)
	return tasks, err
}

func (r *repository) UpdateTaskStatus(ctx context.Context, taskID, status string) error {
	query := `UPDATE tasks SET status = $1 WHERE id = $2`
	_, err := r.dbPool.Exec(ctx, query, status, taskID)
	return err
}

func (r *repository) GetAllTasks(ctx context.Context, filter map[string]string, sort string) ([]domain.Task, error) {
	query := `SELECT * FROM tasks WHERE 1=1`
	args := []interface{}{}
	i := 1
	for k, v := range filter {
		query += fmt.Sprintf(" AND %s = $%d", k, i)
		args = append(args, v)
		i++
	}
	if sort != "" {
		query += fmt.Sprintf(" ORDER BY %s", sort)
	}
	var tasks []domain.Task
	err := pgxscan.Select(ctx, r.dbPool, &tasks, query, args...)
	return tasks, err
}

func (r *repository) GetTaskSummary(ctx context.Context) ([]domain.TaskSummary, error) {
	query := `SELECT assignee_id, COUNT(*) as total_tasks, SUM(CASE WHEN status = 'Completed' THEN 1 ELSE 0 END) as completed_tasks FROM tasks GROUP BY assignee_id`
	var summaries []domain.TaskSummary
	err := pgxscan.Select(ctx, r.dbPool, &summaries, query)
	return summaries, err
}

package taskhdl

import (
	"net/http"

	"kn-assignment/internal/constant"
	errors "kn-assignment/internal/core/error"
	"kn-assignment/internal/handler/dto"
	"kn-assignment/internal/log"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body dto.CreateTaskRequest true "Task"
// @Success 201 {object} domain.Task
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks [post]
// @example json { "title": "New Task", "description": "Task description", "due_date": "2024-12-31T23:59:59Z" }
func (h *handler) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()

	var task dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Errorf(ctx, "error binding task: %v", err)
		c.JSON(http.StatusBadRequest, errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Invalid request payload"))
		return
	}
	userId := c.GetString("userId")

	if err := h.svc.CreateTask(ctx, task.ToDomain(), userId); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, task)
}

// @Summary Assign a task to an employee
// @Description Assign a task to an employee
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Param assigneeID body dto.AssignTaskRequest true "Assignee ID"
// @Success 200
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks/{taskID}/assign [patch]
func (h *handler) AssignTask(c *gin.Context) {
	ctx := c.Request.Context()

	var assignee dto.AssignTaskRequest
	if err := c.ShouldBindJSON(&assignee); err != nil {
		log.Errorf(ctx, "error binding assignee: %v", err)
		c.JSON(http.StatusBadRequest, errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Invalid request payload"))
		return
	}

	taskID := c.Param("taskID")

	if err := h.svc.AssignTask(ctx, taskID, assignee.AssigneeID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dto.BaseResponse{Message: "Task assigned successfully"})
}

// @Summary Get tasks by assignee
// @Description Get tasks assigned to a specific user
// @Tags tasks
// @Produce json
// @Param assigneeID path string true "Assignee ID"
// @Success 200 {array} domain.Task
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks/assignee/{assigneeID} [get]
func (h *handler) GetTasksByAssignee(c *gin.Context) {
	assigneeID := c.Param("assigneeID")
	tasks, err := h.svc.GetTasksByAssignee(c.Request.Context(), assigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary Update task status
// @Description Update the status of a specific task
// @Tags tasks
// @Accept json
// @Produce json
// @Param taskID path string true "Task ID"
// @Param status body dto.UpdateTaskStatusRequest true "Status"
// @Success 200 body dto.BaseResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks/{taskID}/status [patch]
func (h *handler) UpdateTaskStatus(c *gin.Context) {
	ctx := c.Request.Context()

	taskID := c.Param("taskID")
	userId := c.GetString("userId")

	isAssigned, err := h.svc.VerifyTaskAssignment(c.Request.Context(), taskID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !isAssigned {
		c.JSON(http.StatusForbidden, errors.NewCustomErrorWithMessage(constant.ErrCodeForbidden, "You can only update tasks assigned to you"))
		return
	}

	var status dto.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&status); err != nil {
		log.Errorf(ctx, "error binding status: %v", err)
		c.JSON(http.StatusBadRequest, errors.NewCustomErrorWithMessage(constant.ErrCodeInvalidRequest, "Invalid request payload"))
		return
	}
	if err := h.svc.UpdateTaskStatus(ctx, taskID, status.Status, userId); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dto.BaseResponse{Message: "Task status updated successfully"})
}

// @Summary Get all tasks
// @Description Get all tasks with optional filtering and sorting
// @Tags tasks
// @Produce json
// @Param assignee query string false "Assignee ID"
// @Param status query string false "Status"
// @Param sort query string false "Sort"
// @Success 200 {array} domain.Task
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks [get]
func (h *handler) GetAllTasks(c *gin.Context) {
	userRole := c.GetString("role")
	userID := c.GetString("userId")
	filter := map[string]string{}
	if assignee := c.Query("assignee"); assignee != "" {
		filter["assignee_id"] = assignee
	}
	if status := c.Query("status"); status != "" {
		filter["status"] = status
	}
	sort := c.Query("sort")
	tasks, err := h.svc.GetAllTasks(c.Request.Context(), userRole, userID, filter, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary Get task summary
// @Description Get a summary of tasks for each employee
// @Tags tasks
// @Produce json
// @Success 200 {array} domain.TaskSummary
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /tasks/summary [get]
func (h *handler) GetTaskSummary(c *gin.Context) {
	summaries, err := h.svc.GetTaskSummary(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, summaries)
}

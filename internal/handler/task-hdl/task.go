package taskhdl

import (
	"kn-assignment/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body domain.Task true "Task"
// @Success 201 {object} domain.Task
// @Router /tasks [post]
func (h *handler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.svc.CreateTask(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *handler) GetTasksByAssignee(c *gin.Context) {
	assigneeID := c.Param("assigneeID")
	tasks, err := h.svc.GetTasksByAssignee(c.Request.Context(), assigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *handler) UpdateTaskStatus(c *gin.Context) {
	taskID := c.Param("taskID")
	var status struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.UpdateTaskStatus(c.Request.Context(), taskID, status.Status); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

// @Summary Get all tasks
// @Description Get all tasks with optional filtering and sorting
// @Tags tasks
// @Produce json
// @Param assignee query string false "Assignee ID"
// @Param status query string false "Status"
// @Param sort query string false "Sort"
// @Success 200 {array} domain.Task
// @Router /tasks [get]
func (h *handler) GetAllTasks(c *gin.Context) {
	filter := map[string]string{}
	if assignee := c.Query("assignee"); assignee != "" {
		filter["assignee_id"] = assignee
	}
	if status := c.Query("status"); status != "" {
		filter["status"] = status
	}
	sort := c.Query("sort")
	tasks, err := h.svc.GetAllTasks(c.Request.Context(), filter, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *handler) GetTaskSummary(c *gin.Context) {
	summaries, err := h.svc.GetTaskSummary(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, summaries)
}

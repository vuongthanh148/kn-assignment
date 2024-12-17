package router

import (
	"kn-assignment/docs"
	taskhdl "kn-assignment/internal/handler/task-hdl"
	"kn-assignment/property"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HandlerList struct {
	TaskHandler taskhdl.Handler
}

const serviceBaseURL = "/api/v1"

func InitRouter(app *gin.Engine, h HandlerList) {
	docs.SwaggerInfo.Title = property.Get().Server.ServiceName
	docs.SwaggerInfo.Description = property.Get().Server.ServiceDescription
	docs.SwaggerInfo.Version = property.Get().Server.ApiDocsVersion
	docs.SwaggerInfo.Host = property.Get().Server.ApiDocsHost
	docs.SwaggerInfo.Schemes = []string{property.Get().Server.ApiDocsSchema}
	docs.SwaggerInfo.BasePath = serviceBaseURL

	v1 := app.Group("/api/v1")
	registerBasePath(v1, h)

	swaggerV1 := app.Group(serviceBaseURL + "/v1")
	registerBasePath(swaggerV1, h)

	// Add API doc endpoint to router
	if property.Get().Server.ApiDocs {
		docPath := "/docs"
		app.GET(docPath, func(ctx *gin.Context) { ctx.Redirect(http.StatusTemporaryRedirect, docPath+"/swagger/index.html") })
		app.GET(docPath+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func registerBasePath(v1 *gin.RouterGroup, h HandlerList) {
	v1.POST("/tasks", h.TaskHandler.CreateTask)
	v1.GET("/tasks/assignee/:assigneeID", h.TaskHandler.GetTasksByAssignee)
	v1.PATCH("/tasks/:taskID/status", h.TaskHandler.UpdateTaskStatus)
	v1.GET("/tasks", h.TaskHandler.GetAllTasks)
	v1.GET("/tasks/summary", h.TaskHandler.GetTaskSummary)
}

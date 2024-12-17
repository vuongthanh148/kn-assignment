package main

import (
	"context"
	"kn-assignment/infrastructure"
	tasksvc "kn-assignment/internal/core/service/task-svc"
	taskhdl "kn-assignment/internal/handler/task-hdl"
	taskrepo "kn-assignment/internal/repository/postgres/task-repo"
	"kn-assignment/internal/router"
	"kn-assignment/middleware"
	"kn-assignment/property"
	"kn-assignment/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	property.Init(ctx)

	// init infrastructure
	pgx, scanapi := infrastructure.NewPostgres(ctx)
	flavor := infrastructure.NewQueryBuilder()

	// init repository
	taskRepository := taskrepo.New(pgx, scanapi, flavor)

	// init service
	taskService := tasksvc.New(taskRepository)

	// init handler
	taskHandler := taskhdl.New(taskService)

	// init server
	engine := server.InitServer()

	engine.Use(middleware.RequestLogger(ctx))
	engine.Use(middleware.ResponseLogger(ctx))

	// init router
	route := router.HandlerList{
		TaskHandler: taskHandler,
	}

	router.InitRouter(engine, route)

	serverHost := property.Get().Server.Host
	serverPort := property.Get().Server.Port

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	server.StartServerWithCtx(ctx, engine, serverHost, serverPort)

	// Wait for a termination signal
	sig := <-gracefulStop
	log.Printf("Received signal: %v", sig)

	// Give some time for the cron job to stop
	time.Sleep(2 * time.Second)
	log.Print("Application shutdown completed")
}

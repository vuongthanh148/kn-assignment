package main

import (
	"context"
	"kn-assignment/infrastructure"
	authsvc "kn-assignment/internal/core/service/auth-svc"
	tasksvc "kn-assignment/internal/core/service/task-svc"
	authhdl "kn-assignment/internal/handler/auth-hdl"
	taskhdl "kn-assignment/internal/handler/task-hdl"
	"kn-assignment/internal/log"
	authrepo "kn-assignment/internal/repository/postgres/auth-repo"
	taskrepo "kn-assignment/internal/repository/postgres/task-repo"
	userrepo "kn-assignment/internal/repository/postgres/user-repo"
	"kn-assignment/internal/router"
	"kn-assignment/middleware"
	"kn-assignment/property"
	"kn-assignment/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: \"Authorization: Bearer {token}\""
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	property.Init(ctx)

	// Run database migrations
	postgres := property.Get().Postgres
	databaseUrl := "postgres://" + postgres.User + ":" + postgres.Password + "@" + postgres.Host + ":" + postgres.Port + "/" + postgres.Database + "?sslmode=disable"
	server.RunMigrations(ctx, databaseUrl)

	// init infrastructure
	pgx, scanapi := infrastructure.NewPostgres(ctx)
	flavor := infrastructure.NewQueryBuilder()

	// init repository
	taskRepository := taskrepo.New(pgx, scanapi, flavor)
	authRepository := authrepo.New(pgx, scanapi, flavor)
	userRepository := userrepo.New(pgx, scanapi, flavor)

	// init service
	taskService := tasksvc.New(taskRepository, userRepository)
	authService := authsvc.New(authRepository)

	// init handler
	taskHandler := taskhdl.New(taskService)
	authHandler := authhdl.New(authService)

	// init server
	engine := server.InitServer()

	engine.Use(middleware.RequestLogger(ctx))
	engine.Use(middleware.ResponseLogger(ctx))

	// init router
	route := router.HandlerList{
		TaskHandler: taskHandler,
		AuthHandler: authHandler,
	}

	router.InitRouter(engine, route)

	serverHost := property.Get().Server.Host
	serverPort := property.Get().Server.Port

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	server.StartServerWithCtx(ctx, engine, serverHost, serverPort)

	// Wait for a termination signal
	sig := <-gracefulStop
	log.Infof(ctx, "Received signal: %v", sig)

	// Give some time for the cron job to stop
	time.Sleep(2 * time.Second)
	log.Info(ctx, "Application shutdown completed")
}

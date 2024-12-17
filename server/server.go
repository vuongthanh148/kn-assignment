package server

import (
	"context"
	"fmt"
	"kn-assignment/property"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	mode := strings.TrimSpace(property.Get().Gin.Mode)
	switch mode {
	case "release", "debug", "test":
	default:
		mode = "release"
	}
	gin.SetMode(mode)
	return gin.New()
}

func StartServerWithCtx(ctx context.Context, h http.Handler, host string, port string) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: h,
	}

	go func() {
		log.Printf("server running at: %s:%s", host, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error servier listen and serve: %v\n", err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-gracefulStop

	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shut down server error: %v\n", err)
	}
	select {
	case <-ctx.Done():
		log.Print("timeout of 10 seconds.\n")
	default:
	}
	log.Print("Server exiting")
}

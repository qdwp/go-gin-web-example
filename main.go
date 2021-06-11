package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/log"
	"example.com/route"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"
)

// Initialize resource or load configrations.
func initResources() {
	log.InitLogger()
}

func main() {
	env := os.Getenv("GIN_MODE")
	if "release" == env {
		gin.SetMode(gin.ReleaseMode)
	}

	// initialize resources
	initResources()

	routerEngine := route.GetRouterEngine()
	srv := &http.Server{
		Addr:    ":80",
		Handler: routerEngine,
	}

	log.Infof("Start http server ... %v", srv.Addr)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

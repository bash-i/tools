package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	signalChan chan os.Signal
	httpServer *http.Server
)

func newRouter() *gin.Engine {
	if conf.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	engine := gin.Default()

	engine.POST(webhook, xrayWebhookHandler)
	// TODO: AddRouter

	return engine
}

func main() {
	router := newRouter()
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	signalChan = make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	logger.Println("Get Signal:", sig)

	logger.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logger.Println("Server exiting")
}

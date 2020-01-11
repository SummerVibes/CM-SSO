package main

import (
	"CM-SSO/common/route"
	"context"
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	. "github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func readConfig()  {
	env := *flag.String("env","dev","环境变量")
	log.Printf("current env is: %s",env)
	path := *flag.String("start","./","配置文件路径")
	AddConfigPath(path)
	if env == "prod"{
		SetConfigName("config-prod")
	}else {
		SetConfigName("config-dev")
	}
	err := ReadInConfig()
	if err != nil { // Handle errors reading the start file
		log.Fatalf("read config error: %s \n", err)
	}
	log.Println("read config success")
	//检测配置文件变化
	WatchConfig()
	OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config changed: %s", e.Name)
	})
}

func configureLog()  {
	//定义路由日志格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("request %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}



func startServer(){
	router := route.GetRoutes()
	srv:= &http.Server{
		Addr:   GetString("app.host"),
		Handler: router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//启动服务器
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server start failed: %s\n", err)
		}
	}()
	log.Printf("server start success: %s",srv.Addr)

	//优雅停机
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("shutdown error", err)
	}
	log.Println("server has been closed")
}
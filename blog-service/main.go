package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lvdbing/books/blog-service/global"

	"github.com/lvdbing/books/blog-service/internal/model"
	"github.com/lvdbing/books/blog-service/internal/routers"
	"github.com/lvdbing/books/blog-service/pkg/logger"
	"github.com/lvdbing/books/blog-service/pkg/setting"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @Title 博客系统
// @version 1.0
// @description Go语言编写博客后台
// @termsOfService blabla
func main() {
	global.Logger.Infof("%s: %s logger is working...", "Hi", "hoho")

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %v", err)
	}

	setupLogger()

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine error: %v", err)
	}
}

func setupSetting() error {
	configFiles := []setting.ConfigFile{
		{Name: "config", Path: "configs/", Type: "yaml"},
		{Name: "password", Path: "configs/", Type: "yaml"},
	}
	setting, err := setting.NewSetting(configFiles)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupLogger() {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lvdbing/books/blog-service/global"

	"github.com/lvdbing/books/blog-service/internal/model"
	"github.com/lvdbing/books/blog-service/internal/routers"
	"github.com/lvdbing/books/blog-service/pkg/logger"
	"github.com/lvdbing/books/blog-service/pkg/setting"
	"github.com/lvdbing/books/blog-service/pkg/tracer"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// @Title 博客系统
// @version 1.0
// @description Go语言编写博客后台
// @termsOfService blabla
func main() {
	// go build -ldflags "-X main.buildVersion=1.0.1"
	// .\blog-service.exe -version
	if isVersion {
		fmt.Printf("build_version: %s\n", buildVersion)
		return
	}

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func init() {
	setupFlag()
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %v", err)
	}

	setupLogger()

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine error: %v", err)
	}

	err = setUpTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

var (
	isVersion    bool
	buildVersion string
)

func setupFlag() {
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
}

type settingErr struct {
	setting *setting.Setting
	err     error
}

func (s *settingErr) readSection(k string, v interface{}) {
	if s.err != nil {
		return
	}
	s.err = s.setting.ReadSection(k, v)
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
	se := &settingErr{setting: setting}
	se.readSection("Server", &global.ServerSetting)
	se.readSection("App", &global.AppSetting)
	se.readSection("Database", &global.DatabaseSetting)
	se.readSection("JWT", &global.JWTSetting)
	se.readSection("Email", &global.EmailSetting)
	if se.err != nil {
		return se.err
	}

	// err = setting.ReadSection("Server", &global.ServerSetting)
	// if err != nil {
	// 	return err
	// }
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	// err = setting.ReadSection("App", &global.AppSetting)
	// if err != nil {
	// 	return err
	// }

	// err = setting.ReadSection("Database", &global.DatabaseSetting)
	// if err != nil {
	// 	return err
	// }

	// err = setting.ReadSection("JWT", &global.JWTSetting)
	// if err != nil {
	// 	return err
	// }
	global.JWTSetting.Expire *= time.Second

	// err = setting.ReadSection("Email", &global.EmailSetting)
	// if err != nil {
	// 	return err
	// }

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

func setUpTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

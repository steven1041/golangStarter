package main

import (
	"context"
	"flag"
	"fmt"
	"golangStarter/controller"
	"golangStarter/pkg/snowflake"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"golangStarter/routes"

	"golangStarter/dao/redis"

	"golangStarter/dao/mysql"

	"golangStarter/logger"

	"golangStarter/settings"
)

// Golang Web开发通用的脚手架
var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "cfg", "./config.yaml", "config file path")
	flag.Parse()
	//1.加载配置
	if err := settings.Init(configPath); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	//3.初始化MySql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err:%v\n", err)
		return
	}
	//5.注册路由
	r := routes.SetUp()
	//6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	//kill默认会发送syscall.SIGTERM信号
	//kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发这个信号
	//kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	//signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server...")
	//创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//5秒内优雅的关闭服务(将没有处理的请求处理完再关闭服务),超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}

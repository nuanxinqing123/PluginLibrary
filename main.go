// -*- coding: utf-8 -*-
// @Time    : 2022/5/12 14:51
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : main.go

package main

import (
	"PluginLibrary/dataSource"
	"PluginLibrary/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// 加载配置文件
	conf := dataSource.LoadConfig()

	// 注册路由
	r := routes.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: r,
	}

	fmt.Println(" ")
	fmt.Println("服务监听端口:" + strconv.Itoa(conf.Port))
	fmt.Println(" ")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置5秒超时
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	fmt.Println("Service ready to shut down")

	// 创建五秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 五秒内优雅关闭服务（将未处理完成的请求处理完再关闭服务），超过五秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Service timed out has been shut down：", err.Error())
	}
	fmt.Println("Service has been shut down")
}

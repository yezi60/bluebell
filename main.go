package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bluebell/controllers"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"

	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routers"
	"bluebell/settings"

	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// @title bluebell
// @version 1.0
// @description 博客论坛
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url https://soleaf.xyz
// @contact.email yezidexiatian60@sina.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath /api/v1
func main() {

	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: config.yaml")
	//	return
	//}

	// 1. 加载配置文件
	if err := settings.Init("conf/config.yaml"); err != nil {
		fmt.Printf("init file failed: %v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed: %v\n", err)
		return
	}

	defer zap.L().Sync()
	zap.L().Debug("logger init successs")

	// 3. 初始化Mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed: %v\n", err)
		return
	}
	defer redis.Close()

	// 5. 雪花算法生成id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed,err :%v\n", err)
		return
	}

	// 参数校验模块的err翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err :%v\n", err)
		return
	}

	// 5. 注册路由
	r := routers.SetupRouter(settings.Conf.Mode)

	fmt.Println(settings.Conf.Port)
	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: \n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

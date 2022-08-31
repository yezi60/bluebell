package routers

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"
	"time"

	_ "bluebell/docs"

	"github.com/gin-contrib/pprof"
	_ "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		//设置为发行模式
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// 日志，恢复，限流
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("templates/index.html")
	r.Static("/static", "./static")

	// 前端服务
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登陆业务路由
	v1.POST("/login", controllers.LoginHandler)

	// 根据时间或分数获取帖子列表
	v1.GET("/posts2", controllers.GetPostListHandler2)

	// 社区获取 get请求一般不鉴权
	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailHandler)
	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	v1.GET("/posts", controllers.GetPostListHandler)

	// 应用中间件
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.ReteLimitMiddleware(2*time.Second, 1))
	{

		v1.POST("/post", controllers.CreatePostHandler)

		// 投票
		v1.POST("/vote", controllers.PostVoteController)
	}

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}

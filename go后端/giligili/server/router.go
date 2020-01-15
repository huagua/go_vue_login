package server

import (
	"os"
	"giligili/api"
	"giligili/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改，为什么不能改？
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))		//用来保持一个状态
	r.Use(middleware.Cors())						//解决跨域问题
	r.Use(middleware.CurrentUser())					//验证登录

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}

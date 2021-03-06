package initialize

import (
	"net/http"
	"time"

	"gin-basice/router"

	"gin-basice/middleware"

	"gin-basice/global"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 为用户头像和文件提供静态地址
	Router.StaticFS(global.AdpConfig.Local.Path, http.Dir(global.AdpConfig.Local.Path))
	// https
	// Router.Use(middleware.LoadTls())
	global.AdpLog.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.AdpLog.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.AdpLog.Info("register swagger handler")
	// 超时处理
	Router.Use(middleware.TimeoutMiddleware(time.Second * 60))
	// 处理panic
	Router.Use(middleware.GinRecovery(true))
	// 处理日志
	Router.Use(middleware.DefaultLogger())
	// 限制ip 简单方式
	Router.Use(middleware.IpVerifyMiddleware())
	// 限制ip 复杂方式
	//Router.Use(middleware.DefaultLimit())
	// 反向代理
	//Router.Any("/proxy/*url", func(c *gin.Context) {
	//	uri := c.Param("url")
	//	if uri == "" {
	//		c.JSON(http.StatusNotFound, "")
	//		return
	//	}
	//
	//	var proxyUrl = new(url.URL)
	//	proxyUrl.Scheme = global.FMC_CONFIG.Proxy.Scheme
	//	proxyUrl.Host = global.FMC_CONFIG.Proxy.Host
	//
	//	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	//	tr := &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	}
	//	proxy.Transport = tr
	//	c.Request.URL.Path = uri
	//	c.Request.Host = global.FMC_CONFIG.Proxy.Host
	//	c.Request.RequestURI = uri
	//	c.Request.Proto = global.FMC_CONFIG.Proxy.Scheme
	//
	//	proxy.ServeHTTP(c.Writer, c.Request)
	//})

	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	V1RouterGroup := Router.Group("v1")
	{
		PublicGroup := V1RouterGroup.Group("")
		{
			PublicGroup.GET("/ping", func(c *gin.Context) {
				global.AdpLog.Info("pong")
				c.JSON(200, "pong")
			})
		}
		{
			systemRouter.InitBaseRouter(PublicGroup)
		}
		PrivateGroup := V1RouterGroup.Group("")
		PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
		{
			systemRouter.InitCasbinRouter(PrivateGroup)
		}
	}

	//InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.AdpLog.Info("router register success")
	return Router
}

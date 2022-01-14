package routes

import (
	"gin-example/common"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var (
	Router *gin.Engine
	routes = []Option{}
)

// 将各个路由加入路由列表
func include(opts ...Option) {
	routes = append(routes, opts...)
}

func initMiddleware() error {
	common.IsExistFolder("log")
	file, err := os.OpenFile("log/"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	// 设置默认的日志输出，此处为文件以及标准输出
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	// 当程序崩溃错误时，写入的方式
	// gin.DefaultErrorWriter = io.MultiWriter(errFile)
	return nil
}

func init() {

	if err := initMiddleware(); err != nil {
		log.Fatalln(err)
	}

	Router := gin.Default()
	include(basicRoute, bindRoute, goroutineRoute, uploadRoute, panicRoute)

	// 注册路由
	for _, r := range routes {
		r(Router)
	}

	// 404 路由,如果不设置，则默认返回 404状态
	// Router.NoRoute(redirect)
}

/*
重定向
	c.Redirect(code int, url string) 主要用于外部跳转；内部跳转不好传递主体
	c.Request.URL.Path = url string, router.HandleContext(c) 主要用于内部跳转
*/
func redirect(c *gin.Context) {
	// c.Redirect(http.StatusOK, "http://www.baidu.com")
	// c.Redirect(http.StatusMovedPermanently, "/response")
	c.Request.URL.Path = "/404"
	// router.HandleContext(c)
}

// 404 路由
func noRoute(c *gin.Context) {
	c.String(http.StatusOK, "404")
}

/*
使用中间件
	router.Use(...gin.HandlerFunc) 定义全局中间件
	router.Group(uri string, ...gin.HandlerFunc) 路由组中间件
	router.Method(uri string, ...gin.HandlerFunc) 路由中间件
*/
func useMiddleware() {
	// router.Use(middleware())
	// router.Group(uri string, middleware())
	// router.GET(uri string, middleware(), handleFunc)
}

// Recovery 中间件接管任何 panic 恢复，如果出现 panic，则写入并返回一个 500 的错误
// Logger 中间件用于将日志写入 gin.DefaultWriter 中
func middleware() {
	// Defalut 默认使用了 Logger、Recovery 中间件
	// gin.Default()

	// Router = gin.New()
	// // 重写 Logger
	// Router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s---%s\n", params.Request.URL, params.TimeStamp.Format("2006-01-02"))
	// }))

	// Router.Use(gin.Recovery())
}

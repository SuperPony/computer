package routes

import (
	"gin-example/app/basic"
	"gin-example/app/bind"
	"gin-example/app/goroutine"

	"github.com/gin-gonic/gin"
)

var basicRoutes = basic.Routes{}
var bindRoutes = bind.Routes{}
var goroutineRoutes = goroutine.Routes{}

func basicRoute(e *gin.Engine) {
	params := e.Group("/params")
	{
		params.GET("get/:name", basicRoutes.GetParams)
		params.POST("/post", basicRoutes.PostParams)
	}
}

func bindRoute(e *gin.Engine) {
	bind := e.Group("/bind")
	{
		bind.GET("header", bindRoutes.ShouldBindHeader)
		bind.GET("uri", bindRoutes.ShouldBindURI)
		bind.POST("should_bind", bindRoutes.ShouldBind)
	}
}

func goroutineRoute(e *gin.Engine) {
	e.GET("/goroutine", goroutineRoutes.Demo)
}

func uploadRoute(e *gin.Engine) {
	upload := e.Group("/upload")
	{
		upload.POST("single", basicRoutes.Upload)
		upload.POST("all", basicRoutes.UploadAll)
	}
}

func panicRoute(e *gin.Engine) {
	e.GET("/panic", basicRoutes.Panic)
}

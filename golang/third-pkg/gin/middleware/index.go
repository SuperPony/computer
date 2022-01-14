package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件: func name() gin.HandlerFunc
// 	c.Set(key string, value interface{}) 为上下文定义键值对
// 	c.Get(key string) (value interface{}, isexts bool) 返回指定的上下文值
// 	c.MustGet(key string) value interface{} 返回指定的上下文值，该键值对必须存在，否则触发 panics
// 	c.Next() 进入请求，当前中间件是前置或后置，完全取决于业务代码写在 c.Next() 之前或之后
func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().Unix()
		c.Set("time", t)
		// 请求前
		c.Next()
		// 请求后
		// some code
	}
}

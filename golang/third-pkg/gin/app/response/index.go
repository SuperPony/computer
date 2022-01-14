package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

/*
响应输出
	c.String(code int, format string, values ...interface{}) 字符串
	c.Json(code int, obj interface{})
	c.AsciiJSON(code int, obj interface{}) 生成仅有 ASCII 字符的 json，非 ASCII 字符被转移
	c.PureJSON(code int, obj interface{})  生成原样输出的 HTML，HTML 中的特殊字符不会被转义
	c.XMl(code int, obj interface{})
	c.YAML(code int, obj interface{})
*/
func (r *Routes) Response(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": c.Query("name"),
	})

	// c.PureJSON(code int, obj interface{})
	// c.AsciiJSON(code int, obj interface{})
	// c.String(200, "some string")
	// c.XML(code int, obj interface{})
	// c.YAML(code int, obj interface{})
}

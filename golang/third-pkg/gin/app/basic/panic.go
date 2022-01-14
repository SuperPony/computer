package basic

import "github.com/gin-gonic/gin"

func (r *Routes) Panic(c *gin.Context) {
	panic("is error")
}

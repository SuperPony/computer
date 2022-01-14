package goroutine

import (
	"gin-example/common"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

var wg sync.WaitGroup

// 在中间件或处理程序中使用 goroutine 时，必须使用上下文的只读副本，而不是上下文。
//	c.Copy() *gin.Context
func (r *Routes) Demo(c *gin.Context) {
	cCp := c.Copy()
	common.IsExistFolder("log")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeLog(cCp)
	}

	wg.Wait()
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func writeLog(c *gin.Context) {
	file, _ := os.OpenFile("log/goroutine.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer func() {
		file.Close()
		wg.Done()
	}()

	logger := log.New(file, "", log.LstdFlags)
	logger.Printf("[%s]  %s?%s", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)
}

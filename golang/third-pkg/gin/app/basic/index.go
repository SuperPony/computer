package basic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
}

type Data struct {
	Name string `json:"name,omitempty" db:"name" form:"name" binding:"required"`
	Sex  string `json:"sex,omitempty" db:"sex" form:"sex"`
	Age  string `json:"age,omitempty" db:"age" form:"age"`
}

// 获取 get 参数。
// 	c.DefaultQuery(key, defaultVal string) 获取 Query 参数，带有默认值
//	c.DefaultQuery(key string) Query 参数
//	c.Param(key string) API 参数
func (r *Routes) GetParams(c *gin.Context) {
	age := c.DefaultQuery("age", "33")
	sex := c.Query("sex")
	// API 参数
	name := c.Param("name")
	fmt.Printf("DefaultQuery 具有默认值:%v, Query 不具有默认值：%v, api 参数：%v \n", age, sex, name)
}

func (r *Routes) PostParams(c *gin.Context) {
	d := Data{
		Name: c.PostForm("name"),
		Sex:  c.DefaultPostForm("sex", "1"),
		Age:  c.DefaultPostForm("age", "36"),
	}

	fmt.Println(d)
	c.JSON(200, d)
}

func (r *Routes) middleware_demo(c *gin.Context) {
	time, ok := c.Get("time")
	if !ok {
		time = nil
	}
	c.JSON(http.StatusOK, gin.H{
		"request_time": time,
	})
}

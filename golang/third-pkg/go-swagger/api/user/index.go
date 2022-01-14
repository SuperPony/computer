package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取指定用户
type UserItem struct {
	// 用户 ID
	// in: path
	// required: true
	Id int `json:"id" uri:"id"`
}

func Get(c *gin.Context) {

	user := &UserItem{}
	if err := c.BindUri(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is empty",
			"code":    1001,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// 添加用户
type CreateUser struct {
	// 用户 ID
	// required: true
	// min: 1
	// example: 123
	ID int `json:"id"`

	// 用户姓名
	// required: true
	// min length: 3
	// max length: 10
	// example: 张三
	Name string `json:"name"`

	// 用户性别，1男、2女、3未知
	// required: true
	// min: 1
	// max: 3
	// example: 1
	Sex uint `json:"sex"`
}

func Create(c *gin.Context) {

	user := &CreateUser{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    1001,
		})
		return
	}

	// do some thing

	c.JSON(http.StatusOK, user)
}

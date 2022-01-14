// 模型绑定到结构体时，要注意字段的注解是否和绑定的类型一致
// Json -> Field:`json:"field"`
// xml -> Field:`xml:"field"`
// form -> Field:`form:"field"`
// Header -> Field:`header:"field"`
// Uri -> Field:`uri:"field"`

package bind

import (
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

/*
模型绑定：用于将请求的主体直接与类型进行绑定，可用于 query、json、form、xml；在结构体中，字段如果设置了 `binding:"required"` 修饰，则为空时，报错。
	Must bind 系列：当绑定存在错误后，请求将强制被 `c.AbortWithError(400, err).SetType(ErrorTypeBind)` 终止，响应状态码被强制改为 400，后续修改状态码时会被告警，且无法修改，请求头 Content-Type 改为 `text/plain; charset=utf-8`。
	Should bind 系列: 可以更好的控制模型绑定的行为,仅仅返回错误。
*/
func (r *Routes) ShouldBind(c *gin.Context) {
	var d Data

	if err := c.ShouldBind(&d); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		return
	}
	// 根据请求头 Content-Type，自行推断使用那种绑定器。也可用来绑定表单
	// c.Bind(obj)
	// c.ShouldBind(obj)

	// c.ShouldBindJSON(obj) json 绑定器
	// c.ShouldBindQuery(obj) query 绑定器
	// c.ShouldBindXML(obj) xml 绑定器

	c.JSON(200, d)
}

type Header struct {
	Host      string `header:"Host"`
	UserAgent string `header:"User-Agent"`
}

func (r *Routes) ShouldBindHeader(c *gin.Context) {

	h := Header{}
	err := c.ShouldBindHeader(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, h)
}

type Uri struct {
	Name string `uri:"name"`
}

func (r *Routes) ShouldBindURI(c *gin.Context) {

	u := Uri{}
	err := c.ShouldBindUri(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, u)
}

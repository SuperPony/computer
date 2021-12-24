package behavioral

import "fmt"

// 定义接口，规定对象必须具有的方法
type Component interface {
	cooke()
}

// 被装饰者
type Tomatoes struct {
}

func (t *Tomatoes) cooke() {
	fmt.Println("炒西红柿")
}

// 定义装饰器
type WarpTomatoes struct {
	Component
}

// 功能增强
func (c *WarpTomatoes) cooke() {
	fmt.Println("开火、上锅")
	c.Component.cooke()
	fmt.Println("做完饭，我把火关了")
}

func NewWarpTomatoes(c Component) Component {
	return &WarpTomatoes{
		Component: c,
	}
}

func Example() {
	warpTomatoes := NewWarpTomatoes(&Tomatoes{})
	warpTomatoes.cooke()
}

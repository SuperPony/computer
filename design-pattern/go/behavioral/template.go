package structural

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 定义骨架
type CookMenu struct {
}

// 骨架细节的具体实现
func (*CookMenu) fire() {
	// do something...
	fmt.Println("开火")
}

func (*CookMenu) cooke() {
	// do something...
	fmt.Println("做菜")
}

func (*CookMenu) outfire() {
	// do something...
	fmt.Println("关火")
}

// 封装具体步骤
func doCooke(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

// 通过继承的方式，重写细节的具体实现
type Tomatoes struct {
	CookMenu
}

func (t *Tomatoes) cooke() {
	fmt.Println("炒西红柿")
}

type Sushi struct {
	CookMenu
}

func (s *Sushi) fire() {
	fmt.Println("不开火")
}

func (s *Sushi) outfire() {
	fmt.Println("不用关火")
}

func Example() {
	doCooke(&Tomatoes{})
	doCooke(&Sushi{})
}

package creational

import "fmt"

type People struct {
	Name string
	Age  uint
}

type Animal interface {
	Run()
}

func (p *People) Run() {
	fmt.Printf("i'm %s, I'm running \n", p.Name)
}

// 简单工厂，接收一些参数，返回对应实例
// func NewPeople(name string, age uint) *People {
// 	return &People{
// 		Name: name,
// 		Age:  age,
// 	}
// }

// 抽象工厂，返回接口
// func NewPeople(name string, age uint) Animal {
// 	return &People{
// 		Name: name,
// 		Age:  age,
// 	}
// }

// 工厂方法模式,工厂函数参数适合放一些默认的、通用的参数
func NewPeople(age uint) func(name string) *People {
	return func(name string) *People {
		return &People{
			Name: name,
			Age:  age,
		}
	}
}

func Example() {
	newBaby := NewPeople(1)
	_ = newBaby("jack")

	newMan := NewPeople(27)
	_ = newMan("tom")
}

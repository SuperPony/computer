package behavioral

import "fmt"

type Seller interface {
	sell(name string)
}

// 定义目标对象
type station struct {
	stock int
}

func (s *station) sell(name string) {
	if s.stock > 0 {
		s.stock--
		fmt.Printf("火车票售出，购买者 %s, 剩余 %v \n", name, s.stock)
	} else {
		fmt.Println("火车票已售罄")
	}
}

// 定义代理对象
type StationProxy struct {
	station *station
}

// 通过操作代理对象来实现对目标对象的操作
func (s *StationProxy) Sell(name string) {
	// or do something
	if s.station.stock > 0 {
		s.station.stock--
		fmt.Printf("火车票售出，购买者 %s, 剩余 %v \n", name, s.station.stock)
	} else {
		fmt.Println("火车票已售罄")
	}
}

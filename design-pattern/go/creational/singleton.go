package creational

import "sync"

type singleton struct {
}

// 饿汉模式，包被加载时，自动创建
// var ins *singleton = &singleton{}

// func GetInsOr() *singleton {
// 	return ins
// }

var ins *singleton
var mu sync.Mutex
var once sync.Once

// 懒汉模式，并发情况下，需要考虑加锁
// func GetInsOr() *singleton {
// 	if ins == nil {
// 		mu.Lock()
// 		if ins == nil {
// 			ins = &singleton{}
// 		}
// 		mu.Unlock()
// 	}
// 	return ins
// }

// once.Do(f func()) 保证函数只执行一次，因此保证了并发安全，很适合作为用于单例模式
func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

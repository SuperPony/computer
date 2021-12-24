// 选项模式不属于23种设计模式中，而是在 go 中常用的一种设计模式。
// 选项模式适用于实例化一些字段特别多，但是往往只需要部分字段指定设置的场景
// 该模式用于应对 go 语言函数没有默认值的问题

package structural

import (
	"fmt"
	"time"
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = true
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func withTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func withCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

// 经常需要指定设置的字段，可以直接拿到工厂函数中
func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}
	for _, o := range opts {
		o.apply(&options)
	}
	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

func Example() {
	connect, _ := NewConnect("127.0.0.1", withCaching(false))
	fmt.Println(connect) // &{127.0.0.1 false 10}
}

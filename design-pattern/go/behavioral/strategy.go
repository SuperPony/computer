package structural

// 定义策略
type IStrategy interface {
	do(int, int) int
}

// 策略实现 加
type add struct {
}

func (a *add) do(num1, num2 int) int {
	return num1 + num2
}

// 策略实现 减
type reduce struct {
}

func (r *reduce) do(num1, num2 int) int {
	return num1 - num2
}

// 策略执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (o *Operator) setStrategy(strategy IStrategy) {
	o.strategy = strategy
}

// 调用策略
func (o *Operator) culate(num1, num2 int) int {
	return o.strategy.do(num1, num2)
}

// 策略可以随意更换，而不影响 Operator 的所有实现
func Example() {
	operator := &Operator{}

	operator.setStrategy(&add{})
	operator.culate(1, 2)

	operator.setStrategy(&reduce{})
	operator.culate(3, 5)
}

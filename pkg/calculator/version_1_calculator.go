package calculator

type calculator struct {
}

func New() Calculator {
	return &calculator{}
}

func (c *calculator) Add(firstNum int, secondNum int) int {
	return firstNum + secondNum
}
func (c *calculator) Minus(firstNum int, secondNum int) int {
	return firstNum - secondNum
}
func (c *calculator) Divide(firstNum int, secondNum int) int {
	return firstNum / secondNum
}
func (c *calculator) Multiply(firstNum int, secondNum int) int {
	return firstNum * secondNum
}

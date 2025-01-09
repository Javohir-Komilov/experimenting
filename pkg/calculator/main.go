package calculator

type Calculator interface {
	Add(firstNum int, secondNum int) int
	Minus(firstNum int, secondNum int) int
	Divide(firstNum int, secondNum int) int
	Multiply(firstNum int, secondNum int) int
}

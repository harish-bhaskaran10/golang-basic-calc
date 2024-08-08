package basic_calculator

import calc2 "github.com/harish-bhaskaran10/golang-basic-calc2"

func Addition(a int, b int) int {
	return a + b
}

func Subtraction(a int, b int) int {
	return a - b
}

func Mul_in_calc(a int, b int) int {
	return calc2.Multiplication(a, b)
}

func Div_in_calc(a int, b int) int {
	return calc2.Division(a, b)
}

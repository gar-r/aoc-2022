package main

type operator = func(left, right int) int

var add = func(left, right int) int {
	return left + right
}

var subtract = func(left, right int) int {
	return left - right
}

var multiply = func(left, right int) int {
	return left * right
}

var divide = func(left, right int) int {
	return left / right
}

func parseOp(s string) operator {
	switch s {
	case "+":
		return add
	case "-":
		return subtract
	case "*":
		return multiply
	case "/":
		return divide
	default:
		panic(s)
	}
}

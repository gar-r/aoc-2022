package main

import (
	"regexp"
	"strconv"
)

var binaryExprRe = regexp.MustCompile(`([a-z]+) ([+\-*\/]) ([a-z]+)`)

type exprMap map[string]string

func (e exprMap) eval(name string) int {
	expr := e[name]
	if binaryExprRe.MatchString(expr) {
		left, right, op := parseBinary(expr)
		val := parseOp(op)(e.eval(left), e.eval(right))
		e[name] = strconv.Itoa(val)
		return val
	} else {
		n, _ := strconv.Atoi(expr)
		return n
	}
}

func parseBinary(expr string) (left, right, op string) {
	matches := binaryExprRe.FindAllStringSubmatch(expr, -1)
	left = matches[0][1]
	right = matches[0][3]
	op = matches[0][2]
	return
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (e exprMap) solve(name, x string) string {
	if name == x {
		return x
	}
	expr := e[name]
	if binaryExprRe.MatchString(expr) {
		left, right, op := parseBinary(expr)
		a := e.solve(left, x)
		b := e.solve(right, x)
		if !strings.Contains(a, x) && !strings.Contains(b, x) {
			a := mustParseInt(e.solve(left, x))
			b := mustParseInt(e.solve(right, x))
			return strconv.Itoa(parseOp(op)(a, b))
		}
		return fmt.Sprintf("(%s %s %s)", a, op, b)
	} else {
		return expr
	}
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

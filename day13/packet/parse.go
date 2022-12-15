package packet

import (
	"strconv"
	"strings"
)

func Parse(s string) Packet {
	v := &Stack{}
	token := &strings.Builder{}
	for _, c := range s {
		switch c {
		case '[':
			v.Push(nil)
		case ']':
			flush(token, v)
			p := drain(v)
			v.Push(p)
		case ',':
			flush(token, v)
		default:
			token.WriteByte(byte(c))
		}
	}
	return v.Pop()
}

func parseInt(s string) *Int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(s)
	}
	return &Int{i}
}

func drain(v *Stack) *List {
	list := &List{}
	for {
		item := v.Pop()
		if item == nil {
			return list.Reverse()
		} else {
			list.Values = append(list.Values, item)
		}
	}
}

func flush(token *strings.Builder, v *Stack) {
	if token.Len() > 0 {
		v.Push(parseInt(token.String()))
		token.Reset()
	}
}

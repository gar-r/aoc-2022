package packet

import "strings"

type List struct {
	Values []Packet
}

func (l *List) Reverse() *List {
	res := &List{}
	for i := len(l.Values) - 1; i >= 0; i-- {
		res.Values = append(res.Values, l.Values[i])
	}
	return res
}

func (l *List) Compare(other Packet) int {
	switch other.(type) {
	case *List:
		r := other.(*List)
		return l.compareList(r)
	case *Int:
		i := other.(*Int)
		return l.compareList(i.ToList())
	default:
		panic("invalid type")
	}
}

func (l *List) compareList(r *List) int {
	rl := len(r.Values)
	for i := range l.Values {
		if i == rl {
			return 1
		}
		cmp := l.Values[i].Compare(r.Values[i])
		if cmp != 0 {
			return cmp
		}
	}
	if rl == len(l.Values) {
		return 0
	}
	return -1
}

func (l *List) String() string {
	sb := &strings.Builder{}
	sb.WriteString("[")
	for _, item := range l.Values {
		sb.WriteString(item.String())
		sb.WriteString(",")
	}
	sb.WriteString("]")
	return sb.String()
}

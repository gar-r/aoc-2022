package packet

import "fmt"

type Int struct {
	Value int
}

func (i *Int) ToList() *List {
	return &List{
		Values: []Packet{i},
	}
}

func (i *Int) Compare(other Packet) int {
	switch other.(type) {
	case *Int:
		j := other.(*Int)
		return i.compareInt(j)
	case *List:
		return i.ToList().Compare(other)
	default:
		panic("invalid type")
	}
}

func (i *Int) compareInt(j *Int) int {
	if i.Value > j.Value {
		return 1
	}
	if i.Value < j.Value {
		return -1
	}
	return 0
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

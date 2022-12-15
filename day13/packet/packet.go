package packet

import (
	"fmt"
)

type Packet interface {
	fmt.Stringer
	Compare(other Packet) int
}

type Pair struct {
	Left  Packet
	Right Packet
}

func (p *Pair) IsOrdered() bool {
	return p.Left.Compare(p.Right) == -1
}

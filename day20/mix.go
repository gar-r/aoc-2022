package main

func (f *File) Mix() {
	for i := range f.Data {
		f.mix(i)
	}
}

func (f *File) mix(idx int) {
	node := f.Data[idx]
	if node.Value == 0 {
		return
	}
	f.del(node)
	after := f.seek(node, node.Value)
	f.ins(node, after)
}

func (f *File) ins(n, after *Node) {
	n.Next = after.Next
	n.Prev = after
	after.Next.Prev = n
	after.Next = n
}

func (f *File) del(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (f *File) seek(n *Node, steps int) *Node {
	res := n
	steps = steps % (len(f.Data) - 1)
	if steps >= 0 {
		for i := 0; i < steps; i++ {
			res = res.Next
		}
	} else {
		for i := 0; i <= (-steps); i++ {
			res = res.Prev
		}
	}
	return res
}

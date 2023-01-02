package main

import (
	"fmt"
)

type File struct {
	Data []*Node
}

func NewFile(arr []int) *File {
	f := &File{}
	for _, i := range arr {
		f.Data = append(f.Data, &Node{Value: i})
	}
	for i, n := range f.Data {
		if i == len(f.Data)-1 {
			n.Next = f.Data[0]
		} else {
			n.Next = f.Data[i+1]
		}
		if i == 0 {
			n.Prev = f.Data[len(f.Data)-1]
		} else {
			n.Prev = f.Data[i-1]
		}
	}
	return f
}

func (f *File) GetValues() []int {
	res := make([]int, 0)
	start := f.Zero()
	res = append(res, start.Value)
	for node := start.Next; node != start; node = node.Next {
		res = append(res, node.Value)
	}
	return res
}

func (f *File) PrintValues() {
	fmt.Printf("%v\n", f.GetValues())
}

func (f *File) Zero() *Node {
	for _, node := range f.Data {
		if node.Value == 0 {
			return node
		}
	}
	return nil
}

func (f *File) ApplyKey(key int) {
	for _, node := range f.Data {
		node.Value = node.Value * key
	}
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d, n: %d, p: %d", n.Value, n.Next, n.Prev)
}

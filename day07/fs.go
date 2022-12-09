package main

import (
	"strconv"
	"strings"
)

type Composite interface {
	Name() string
	Size() int
}

func parse(s string) Composite {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		panic(s)
	}
	if parts[0] == "dir" {
		return &Directory{
			DirName: parts[1],
		}
	}
	size, _ := strconv.Atoi(parts[0])
	return &File{
		Filename: parts[1],
		Len:      size,
	}
}

type File struct {
	Filename string
	Len      int
}

func (f *File) Name() string {
	return f.Filename
}

func (f *File) Size() int {
	return f.Len
}

type Directory struct {
	DirName  string
	Children []Composite
}

func (d *Directory) Name() string {
	return d.DirName
}

func (d *Directory) Size() int {
	total := 0
	for _, obj := range d.Children {
		total += obj.Size()
	}
	return total
}

func (d *Directory) FindDirectory(name string) *Directory {
	for _, c := range d.Children {
		d, ok := c.(*Directory)
		if ok && d.Name() == name {
			return d
		}
	}
	return nil
}

func (d *Directory) GetDirectories() []*Directory {
	result := make([]*Directory, 0)
	for _, c := range d.Children {
		d, ok := c.(*Directory)
		if ok {
			result = append(result, d)
		}
	}
	return result
}

package main

import (
	"fmt"
	"testing"
)

type People struct {
	Name     string
	Children []People
}

func TestRecurList(t *testing.T) {
	p1 := People{Name: "p1", Children: []People{
		{Name: "c1", Children: []People{{Name: "d1"}}},
		{Name: "c2", Children: []People{{Name: "d2"}}},
	}}
	p2 := People{Name: "p2", Children: []People{
		{Name: "c11", Children: []People{{Name: "d11", Children: []People{{Name: "e11"}}}}},
		{Name: "c21", Children: []People{{Name: "d21"}}},
	}}
	list := []People{p1, p2}

	var fn func([]People) []People
	fn = func(p []People) []People {
		for i, v := range p {
			if v.Name == "d1" || v.Name == "e11" {
				p = append(p[:i], p[i+1:]...)
				continue
			}
			children := fn(v.Children)
			v.Children = append([]People{}, children...)
			p[i] = v
		}
		return p
	}

	res := fn(list)
	fmt.Printf("%+v\n", res)
}

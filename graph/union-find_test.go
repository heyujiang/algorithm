package graph

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUF(10)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	fmt.Println(uf)
	//uf.Union(1,3)
	fmt.Println(uf.Connected(0, 1))
	fmt.Println(uf)

}

package tree

import (
	"fmt"
	"testing"
)

func TestBu(t *testing.T) {
	res := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(preorderTraversal(res))
}

func TestAA(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)

}

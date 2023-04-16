package tree

import (
	"fmt"
	"testing"
)

func TestBu(t *testing.T) {
	res := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(preorderTraversal(res))
}

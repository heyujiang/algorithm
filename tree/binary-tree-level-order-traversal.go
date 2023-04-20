package tree

import (
	"math"
)

//102. 二叉树的层序遍历  &&剑指 Offer 32 - III. 从上到下打印二叉树 I  && 剑指 Offer 32 - III. 从上到下打印二叉树 II
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			n := len(queue)
			level := make([]int, 0)
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
				level = append(level, cur.Val)
			}
			res = append(res, level)
		}
	}

	return res
}

//103. 二叉树的锯齿形层序遍历   && 剑指 Offer 32 - III. 从上到下打印二叉树 III
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		flag := true

		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			n := len(queue)
			level := make([]int, 0)
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
				if flag {
					level = append(level, cur.Val)
				} else {
					level = append([]int{cur.Val}, level...)
				}
			}
			flag = !flag
			res = append(res, level)
		}
	}
	return res
}

//107. 二叉树的层序遍历 II   [自底向上的层序遍历]
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			n := len(queue)
			level := make([]int, 0)
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
				level = append(level, cur.Val)
			}
			res = append([][]int{level}, res...)
		}
	}

	return res
}

//1161. 最大层内元素和
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	maxSum := root.Val
	level := 1
	res := level
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			sum += cur.Val
		}
		if sum > maxSum {
			maxSum = sum
			res = level
		}
		level++
	}

	return res
}

//1302. 层数最深叶子节点的和
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := root.Val
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			sum += cur.Val
		}
		res = sum
	}

	return res
}

//1609. 奇偶树
func isEvenOddTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	flag := true
	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			level = append(level, cur.Val)
		}
		if !checkIsEvenOddTree(level, flag) {
			return false
		}

		flag = !flag
	}
	return true
}

func checkIsEvenOddTree(nums []int, flag bool) bool {
	if flag {
		pre := math.MinInt32
		for _, v := range nums {
			if v%2 == 1 && v > pre {
				pre = v
			} else {
				return false
			}
		}
	} else {
		pre := math.MaxInt32
		for _, v := range nums {
			if v%2 == 0 && v < pre {
				pre = v
			} else {
				return false
			}
		}
	}
	return true
}

//637. 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root != nil {
		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			n := len(queue)
			sum := 0
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
				sum += cur.Val
			}
			res = append(res, float64(sum)/float64(n))
		}
	}

	return res
}

//958. 二叉树的完全性检验
func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	level := float64(0)
	for len(queue) > 0 {
		hasNil := false
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur != nil && hasNil {
				return false
			}
			if cur == nil {
				hasNil = true
			} else {
				queue = append(queue, cur.Left)
				queue = append(queue, cur.Right)
			}
		}
		if n < int(math.Pow(2, level)) && len(queue) > 0 {
			return false
		}
		level++
	}
	return true
}

//919. 完全二叉树插入器
type CBTInserter struct {
	Level     int
	LevelNums [][]*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	levelNums := levelOrderTree(root)
	return CBTInserter{
		Level:     len(levelNums) - 1,
		LevelNums: levelNums,
	}
}

func levelOrderTree(root *TreeNode) [][]*TreeNode {
	res := make([][]*TreeNode, 0)
	if root != nil {
		queue := make([]*TreeNode, 0)
		queue = append(queue, root)

		for len(queue) > 0 {
			n := len(queue)
			level := make([]*TreeNode, 0)
			for i := 0; i < n; i++ {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
				level = append(level, cur)
			}
			res = append(res, level)
		}
	}

	return res
}

func (this *CBTInserter) Insert(val int) int {
	v := &TreeNode{Val: val}
	if len(this.LevelNums[this.Level]) == int(math.Pow(2, float64(this.Level))) {
		this.Level++
		this.LevelNums = append(this.LevelNums, make([]*TreeNode, 0))
	}

	this.LevelNums[this.Level] = append(this.LevelNums[this.Level], v)

	end := len(this.LevelNums[this.Level]) - 1

	parent := this.LevelNums[this.Level-1][end/2]

	if end%2 == 0 {
		parent.Left = v
	} else {
		parent.Right = v
	}

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.LevelNums[0][0]
}

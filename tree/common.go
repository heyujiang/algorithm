package tree

/**
二叉树解题的思维模式分两类 ：
1 . 是否可以通过遍历一遍二叉树得到答案？  如果可以，用一个traverse函数配合外部变量来实现，这叫【遍历】的思维模式。
2 . 是否可以定义一个递归函数，通过子问题（子树）的答案推导愿问题的答案？ 如果可以，写出这个递归函数的定义，并充分利用这个函数的返回值，这叫【分解问题】的思维模式。
*/

/**
无论哪一种思维模式都需要思考 ：
	如果单独抽出一个节点，它都需要做什么事情？需要在什么时候（前/中/后序位置）做？其它节点不用操心，递归函数会帮你在所有节点上执行相同操作。
*/

/**
二叉树的所有问题，就是让你在前中后序位置注入巧妙的代码逻辑，去达到自己的目的，你只需要单独思考每一个节点应该做什么，其它的都不用管，抛给二叉树遍历框架，递归会在所有节点上做相同的操作。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//二叉搜索数  binary-search-tree
//1。 左节点的值都比子节点小，右节点的值都比子节点大。
//2。 每个子树都是二叉搜索数。

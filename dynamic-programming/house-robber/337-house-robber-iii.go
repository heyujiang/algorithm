package house_robber

import (
	"fmt"
	"github.com/heyujiang/algorithm"
	"github.com/heyujiang/algorithm/tree"
)

//337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
/**
从根节点触发器，可以选择偷或者不偷，条件限制了相邻节点不可以同时偷，也就是说如果父节点选择了偷，那么子节点就不能偷，如果父节点选择了不偷，那么子节点可以选择不偷或者偷，因为为了最大值，选择偷与不偷的之间的盗取金额相对大的值
所以状态是不同的节点，条件是偷或者不偷窃，可以用true 、false代替；
状态转移方程是 ：
dp(root,true) = dp(root.Left,false) + root.Right,false) + root.Val
dp(root,false) = max(dp(root.Left,false),dp(root.Left,true)) + max(dp(root.Right,false),dp(root.Right,true))

结果为 ： max(dp(root,true),dp(root,false))
base case 是房屋没有连接的房屋了，即子节点是nil,那么偷盗的财物肯定是0；
*/
func robTree2(root *tree.TreeNode) int {
	memo := make(map[string]int)
	var dp func(root *tree.TreeNode, isSteal bool) int
	dp = func(root *tree.TreeNode, isSteal bool) int {
		if root == nil {
			return 0
		}

		key := fmt.Sprintf("%v%v", root, isSteal)
		if _, ok := memo[key]; ok {
			return memo[key]
		}

		if isSteal {
			memo[key] = dp(root.Left, !isSteal) + dp(root.Right, !isSteal) + root.Val
		} else {
			memo[key] = algorithm.Max(dp(root.Left, isSteal), dp(root.Left, !isSteal)) + algorithm.Max(dp(root.Right, isSteal), dp(root.Right, !isSteal))
		}
		return memo[key]
	}

	return algorithm.Max(dp(root, true), dp(root, false))
}

/**
从跟节点触发器，可以选择偷或者不偷，条件限制了相邻节点不可以同时偷，也就是说如果父节点选择了偷，那么子节点就不能偷，如果父节点选择了不偷，那么子节点可以选择不偷或者偷，因为为了最大值，选择偷与不偷的之间的盗取金额相对大的值

所以如果当前节点选择了偷，那么子节点只能是不偷，可以跳过遍历，直接对子节点的子节点进行计算 + root.Val  ；
如果当前节点选择不偷，那么可以对其子节点进行选择；
*/
func robTree(root *tree.TreeNode) int {
	memo := make(map[*tree.TreeNode]int)
	var dp func(root *tree.TreeNode) int
	dp = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}

		if _, ok := memo[root]; ok {
			return memo[root]
		}

		//偷
		left, right := 0, 0
		if root.Left != nil {
			left = dp(root.Left.Left) + dp(root.Left.Right)
		}
		if root.Right != nil {
			right = dp(root.Right.Left) + dp(root.Right.Right)
		}
		steal := left + right + root.Val

		//不偷
		notSteal := dp(root.Left) + dp(root.Right)

		memo[root] = algorithm.Max(steal, notSteal)

		return memo[root]
	}

	return dp(root)
}

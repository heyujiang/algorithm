package tree

//437.路径总和III (https://leetcode.cn/problems/path-sum-iii/)
//在每条路径上构造前缀和数组，用当前节点的前缀和值减去前面每个节点的前缀和是否等于targetNum ，相等结果加一
func pathSum(root *TreeNode, targetSum int) int {
	var f func(root *TreeNode, nums []int)
	var res = 0
	f = func(root *TreeNode, nums []int) {
		if root == nil {
			return
		}

		nums = append(nums, root.Val+nums[len(nums)-1])
		for i := 0; i < len(nums)-1; i++ {
			if nums[len(nums)-1]-nums[i] == targetSum {
				res++
			}
		}

		f(root.Left, nums)
		f(root.Right, nums)
	}

	f(root, []int{0})
	return res
}

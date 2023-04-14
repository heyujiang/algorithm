package tree

import "math"

//654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var f func(nums []int , left ,right int) *TreeNode
	f = func(nums []int, left, right int) *TreeNode {
		if left < right {
			return nil
		}
		maxNum , index := math.MinInt32,0
		for i := left; i <= right; i++  {
			 if maxNum < nums[i] {
			 	maxNum = nums[i]
			 	index = i
			 }
		}


		return &TreeNode{Val:maxNum,Left:f(nums,left,index-1),Right:f(nums,index+1,right)}
	}

	return f(nums,0,len(nums) - 1)
}


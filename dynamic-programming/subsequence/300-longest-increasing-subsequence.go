package subsequence

//300. 最长递增子序列 (https://leetcode.cn/problems/longest-increasing-subsequence/)
//动态规划解法，时间复杂度O(N^2)
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	//res := 0
	//for _,v := range dp{
	//	if res < v {
	//		res = v
	//	}
	//}
	return res
}

//非动态规划解法，但是时间复杂度更优，O(NlogN)  ；借助蜘蛛纸牌游戏理解的二分查找法：
func lengthOfLISBinarySearch(nums []int) int {
	top := make([]int, len(nums))
	piles := 0 //初始化牌堆为0

	for i := 0; i < len(nums); i++ {
		//搜索左边界的二分查找
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] < nums[i] {
				left = mid + 1
			} else if top[mid] > nums[i] {
				right = mid
			} else if top[mid] == nums[i] {
				right = mid
			}
		}
		//搜索左边界的二分查找

		if left == piles { //没找到合适的牌堆 新建一堆
			piles++
		}

		top[left] = nums[i]
	}
	return piles //牌堆数就是LIS长度
}

package subsequence

//53.最长子数组和 https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	max := make([]int, len(nums))
	max[0] = nums[0]
	res := max[0]
	for i := 1; i < len(nums); i++ {
		max[i] = nums[i]
		if max[i-1]+nums[i] > max[i] {
			max[i] = max[i-1] + nums[i]
		}
		if res < max[i] {
			res = max[i]
		}
	}
	return res
}

func maxSubArrayO1(nums []int) int { //空间复杂度o(1)
	max, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if max+nums[i] > nums[i] {
			max = max + nums[i]
		} else {
			max = nums[i]
		}
		if res < max {
			res = max
		}
	}
	return res
}

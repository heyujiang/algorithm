package array

import "fmt"

//1031. 两个非重叠子数组的最大和(https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays)
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	maxSum := 0
	for i := 0; i <= len(nums)-firstLen; i++ {
		firstSum := sums[i+firstLen] - sums[i]
		for j := i + firstLen; j <= len(nums)-secondLen; j++ {
			secondSum := sums[j+secondLen] - sums[j]
			if firstSum+secondSum > maxSum {
				maxSum = firstSum + secondSum
			}
		}
	}

	for i := 0; i <= len(nums)-secondLen; i++ {
		secondSum := sums[i+secondLen] - sums[i]
		for j := i + secondLen; j <= len(nums)-firstLen; j++ {
			firstSum := sums[j+firstLen] - sums[j]
			if firstSum+secondSum > maxSum {
				maxSum = firstSum + secondSum
				fmt.Println(i, j)
			}
		}
	}

	return maxSum
}

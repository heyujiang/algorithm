package array

//410. 分割数组的最大值
//最大值为下标，分割数组数量为值的数组（递减数组），对这个数组做左边界的二分搜索算法 。
func splitArray(nums []int, k int) int {
	left,right := 0,0
	for _,v := range nums {
		if left < v {
			left = v
		}
		right += v
	}

	for left <= right {
		mid := left + (right - left)/2
		if f(nums,mid) <= k {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	return left
}

func f(nums []int , max int) int {
	m := 0
	i := 0
	for i < len(nums) {
		cap := max
		for cap >= nums[i] {
			cap -= nums[i]
			i++
		}
		m++
	}

	return m
}
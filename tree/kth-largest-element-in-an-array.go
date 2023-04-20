package tree

//215. 数组中的第K个最大元素  && 剑指 Offer II 076. 数组中的第 k 大的数字
func findKthLargest(nums []int, k int) int {
	shuffle(nums)
	return findQuickSort(nums, 0, len(nums)-1, k-1)
}

func findQuickSort(nums []int, lo, hi int, k int) int {
	j := partition(nums, lo, hi)
	if j < k {
		return findQuickSort(nums, j+1, hi, k)
	} else if j > k {
		return findQuickSort(nums, lo, j-1, k)
	} else {
		return nums[j]
	}
}

func partition(nums []int, lo, hi int) int {
	point := nums[lo]

	i, j := lo+1, hi
	for i <= j {
		for i < hi && nums[i] >= point {
			i++
		}
		for j > lo && nums[j] < point {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[j], nums[lo] = nums[lo], nums[j]
	return j
}

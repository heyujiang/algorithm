package array

//34. 在排序数组中查找元素的第一个和最后一个位置   && 剑指 Offer 53 - I. 在排序数组中查找数字 I
func searchRange(nums []int,target int) []int  {
	return []int{searchFirst(nums,target),searchLast(nums,target)}
}

func searchFirst(nums []int, target int) int {
	left,right := 0,len(nums)

	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid+1
		}
	}

	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchLast(nums []int, target int) int {
	left,right := 0,len(nums)

	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == target {
			left = mid + 1
		}else if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid
		}
	}

	if left - 1 < 0 || nums[left - 1] != target {
		return -1
	}

	return left - 1
}
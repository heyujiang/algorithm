package array


//704 . 二分查找
func search(nums []int, target int) int {
	left,right := 0,len(nums)-1

	for left <= right  {
		mid := left + (right - left)/2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			right--
		}else if nums[mid] < target {
			left++
		}
	}
	return -1
}

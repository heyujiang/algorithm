package array

//27 . 移除数组中指定的元素

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow,fast := 0,0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow+1
}

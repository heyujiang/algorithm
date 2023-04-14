package array

//26 .  删除有序数组的重复项
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow ,fast := 0,0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow+1
}

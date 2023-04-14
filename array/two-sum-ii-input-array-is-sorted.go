package array

//167 . 两数之和 ii- 输入有序数组   && 剑指offer 57 . 和为s的两个数字  && 剑指 Offer II 006. 排序数组中两个数字之和

func twoSum(numbers []int,target int) []int {
	left,right := 0,len(numbers) - 1
	for left < right {
		if numbers[left] + numbers[right]  == target {
			return []int{left+1,right+1}
		} else if numbers[left] + numbers[right] < target {
			left++
		}else {
			right--
		}
	}
	return []int{}
}

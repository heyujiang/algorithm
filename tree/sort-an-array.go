package tree

//912 排序数组

var temp []int

func sortArray(nums []int) []int {
	temp = make([]int, len(nums))

	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, left, right int) { //归并排序
	if len(nums) == 1 {
		return
	}

	mid := left + (right-left)/2

	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)

	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}

	i, j := left, mid+1
	for p := left; p <= right; p++ {
		if i == mid+1 {
			nums[p] = temp[j]
			j++
		} else if j == right+1 {
			nums[p] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[p] = temp[j]
			j++
		} else {
			nums[p] = temp[i]
			i++
		}
	}
}

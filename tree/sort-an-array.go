package tree

import "math/rand"

//912 排序数组

var temp []int

func sortArray(nums []int) []int {
	shuffle(nums)
	quickSort(nums, 0, len(nums)-1)

	//temp = make([]int, len(nums))
	//mergeSort(nums, 0, len(nums)-1)
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

func shuffle(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		r := rand.Intn(n - i)
		nums[i], nums[r] = nums[r], nums[i]
	}
}

func quickSort(nums []int, lo, hi int) { //快速排序
	if lo >= hi {
		return
	}
	pivot := nums[lo]
	i, j := lo+1, hi
	for i <= j {
		for i < hi && nums[i] <= pivot {
			i++
		}

		for j > lo && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]

	quickSort(nums, lo, j-1)
	quickSort(nums, j+1, hi)
	return
}

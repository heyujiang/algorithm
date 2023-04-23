package data_structure_design

//monotone stack 单调栈

//96. 下一个更大元素 I （https://leetcode.cn/problems/next-greater-element-i/）
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	tempRes := make(map[int]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 {
			tempRes[nums2[i]] = stack[len(stack)-1]
		} else {
			tempRes[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = tempRes[v]
	}
	return res
}

//96. 下一个更大元素 II （https://leetcode.cn/problems/next-greater-element-ii/）
/*
处理环形数组一般使用%运算符（余数）来模拟环形特效
arr := []int{1, 2, 3, 4, 5}
n, index := len(arr), 0
for {
    // 在环形数组中转圈
    fmt.Println(arr[index%n])
    index++
}
将数组长度翻倍一般就很好处理环形数组的问题
*/

func nextGreaterElementII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0, n)

	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		stack = append(stack, nums[i%n])
	}
	return res
}

/*
处理环形数组一般使用%运算符（余数）来模拟环形特效
将数组长度翻倍一般就很好处理环形数组的问题
*/

//739. 每日温度  (https://leetcode.cn/problems/daily-temperatures/)
//剑指offer II 038. 每日温度 (https://leetcode.cn/problems/iIQa4I/)
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := make([][]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1][0] <= temperatures[i] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 {
			answer[i] = stack[len(stack)-1][1] - i
		} else {
			answer[i] = 0
		}
		stack = append(stack, []int{temperatures[i], i})
	}
	return answer
}

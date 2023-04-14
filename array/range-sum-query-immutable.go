package array

//303. 区域和检索 - 数组不可变
type NumArray struct {
	preNums []int
}


func Constructor(nums []int) NumArray {
	preNums := make([]int,len(nums)+1)
	for i:=0;i<len(nums);i++{
		preNums[i+1] = preNums[i] + nums[i]
	}

	return NumArray{preNums:preNums}
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.preNums[right+1] - this.preNums[left]
}


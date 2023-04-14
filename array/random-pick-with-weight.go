package array

import "math/rand"

//528. 按权重随机选择
func NewSolution(w []int) Solution {
	preNums := make([]int,len(w)+1)
	for i:=0;i<len(w);i++{
		preNums[i+1] = preNums[i] + w[i]
	}
	return Solution{preNums:preNums}
}


type Solution struct {
	preNums []int
}

func (this *Solution) PickIndex() int {
	random := rand.Intn(this.preNums[len(this.preNums)-1]) + 1

	left,right := 0,len(this.preNums) - 1

	for left <= right {
		mid := left + (right - left)/2

		if this.preNums[mid] == random {
			right = mid - 1
		}else if this.preNums[mid] < random {
			left = mid + 1
		}else if this.preNums[mid] > random {
			right = mid - 1
		}
	}

	return left - 1
}
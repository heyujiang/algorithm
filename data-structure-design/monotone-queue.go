package data_structure_design

//monotone queue 单调队列
type MonotoneQueue struct {
	nums []int
}

func (q *MonotoneQueue) Push(x int) {
	for len(q.nums) > 0 && q.nums[len(q.nums)-1] < x {
		q.nums = q.nums[:len(q.nums)-1]
	}
	q.nums = append(q.nums, x)
}

func (q *MonotoneQueue) Pop(x int) {
	if q.Max() == x {
		q.nums = q.nums[1:]
	}
}

func (q *MonotoneQueue) Max() int {
	return q.nums[0]
}

//239. 滑动窗口最大值  (https://leetcode.cn/problems/sliding-window-maximum/)
//剑指 Offer 59 - I. 滑动窗口的最大值 (https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	mq := &MonotoneQueue{nums: make([]int, 0, len(nums))}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			mq.Push(nums[i])
		} else {
			mq.Push(nums[i])
			res = append(res, mq.Max())
			mq.Pop(nums[i-k+1])
		}
	}
	return res
}

//剑指 Offer 59 - II. 队列的最大值  (https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof)
type MaxQueue struct {
	nums          []int
	MonotoneQueue *MonotoneQueue
}

func Constructor() MaxQueue {
	return MaxQueue{nums: make([]int, 0), MonotoneQueue: &MonotoneQueue{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.nums) > 0 {
		return this.MonotoneQueue.Max()
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.nums = append(this.nums, value)
	this.MonotoneQueue.Push(value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.nums) > 0 {
		x := this.nums[0]
		this.nums = this.nums[1:]
		this.MonotoneQueue.Pop(x)
		return x
	}
	return -1

}

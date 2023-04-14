package array

// 1109 . 航班预定统计

func corpFlightBooking(bookings [][]int, n int) []int {
	diff := NewDiffArr(n)

	for _,booking := range bookings {
		diff.Incr(booking[0],booking[1],booking[2])
	}
	return diff.Result()
}

type DiffArr struct {
	diffs []int
}

func NewDiffArr(n int) *DiffArr {
	return &DiffArr{diffs:make([]int,n)}
}

func (d *DiffArr) Incr(start, end  , n int ) {
	d.diffs[start] += n
	if end + 1 < len(d.diffs) {
		d.diffs[end+1] -= n
	}
}

func (d *DiffArr)Result() []int  {
	res := make([]int,len(d.diffs))
	res[0] = d.diffs[0]
	for i:=1;i<len(d.diffs);i++{
		res[i] = d.diffs[i] + res[i-1]
	}
	return res
}



/**
差分数组 ：
nums := make([]int,n)
diff := make([]int,len(nums))
diff[0] = nums[0]
for i:=1;i<len(nums);i++{
	diff[i] = nums[i] - nums[i-1]
}
 */
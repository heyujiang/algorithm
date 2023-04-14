package array

//875 . 爱吃香蕉的珂珂   && 剑指 Offer II 073. 狒狒吃香蕉
//根据速度计算时间，速度越快，时间越短 ； 二分搜索满足时间要求的最小速度；查找左边界的二分搜索算法
func minEatingSpeed(piles []int, h int) int {
	left,right := 1,0

	for _,pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left <= right {
		mid := left + (right - left)/2
		if eatingTime(piles,mid) <= h {
			right = mid - 1
		}else if eatingTime(piles,mid) > h {
			left = mid + 1
		}
	}
	return left
}

func eatingTime(piles []int,x int) int  {
	h := 0
	for _,pile := range piles {
		h += pile/x
		if pile%x != 0 {
			h++
		}
	}
	return h
}



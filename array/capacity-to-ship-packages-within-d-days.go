package array

// 1011. 在 D 天内送达包裹的能力
// 根据运力计算天数 ， 运力增大 ， 天数降低 ， 相当于下标为运力天数为值的递减数组，使用二分搜索搜寻满足要求天数的左边界 ；核心算法 ： 查找左边界的二分搜索  ；
// 下标从运力满足最大包裹开始到可一次运输全部包裹结束； 在递减数组中部分二分搜索
func shipWithinDays(weights []int,days int) int  {
	left ,right := 0,0
	for _,w := range weights {
		if w > left {
			left = w
		}
		right+=w
	}

	for left <= right {
		mid := left+ (right - left) / 2

		if shipWithDaysBy(weights,mid) <= days {
			right = mid - 1
		}else  {
			left = mid + 1
		}
	}

	return left
}

func shipWithDaysBy(weights []int, x int) int {
	days := 0
	i := 0
	for i < len(weights)  {
		cap := x
		for i < len(weights) {
			if cap < weights[i]  {
				break
			}else {
				cap -= weights[i]
				i++
			}
		}
		days++
	}
	return days
}

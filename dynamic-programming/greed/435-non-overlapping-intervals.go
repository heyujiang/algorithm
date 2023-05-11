package greed

import "sort"

//435. 无重叠区间  https://leetcode.cn/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	return n - intervalSchedule(intervals)
}

// 区间调度算法，算出 intvs 中最多有几个互不相交的区间
func intervalSchedule(intvs [][]int) int {
	if len(intvs) == 0 {
		return 0
	}
	// 按 end 升序排序
	sort.Slice(intvs, func(i, j int) bool {
		return intvs[i][1] < intvs[j][1]
	})
	// 至少有一个区间不相交
	count := 1
	// 排序后，第一个区间就是 x
	xEnd := intvs[0][1]
	for _, interval := range intvs {
		start := interval[0]
		if start >= xEnd {
			// 找到下一个选择的区间了
			count++
			xEnd = interval[1]
		}
	}
	return count
}

package greed

import "sort"

//452. 用最少数量的箭引爆气球  https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(intvs [][]int) int {
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
		// 把 >= 改成 > 就行了
		if start > xEnd {
			count++
			xEnd = interval[1]
		}
	}
	return count
}

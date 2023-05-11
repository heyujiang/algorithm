package greed

import (
	"github.com/heyujiang/algorithm"
	"sort"
)

//1024. 视频拼接 https://leetcode.cn/problems/video-stitching/
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	// 按起点升序排列，起点相同的降序排列
	// PS：其实起点相同的不用降序排列也可以，不过我觉得这样更清晰
	sort.Slice(clips, func(i, j int) bool {
		a, b := clips[i], clips[j]
		if a[0] == b[0] {
			return b[1] < a[1]
		}
		return a[0] < b[0]
	})
	// 记录选择的短视频个数
	var res int

	curEnd, nextEnd := 0, 0
	i, n := 0, len(clips)
	for i < n && clips[i][0] <= curEnd {
		// 在第 res 个视频的区间内贪心选择下一个视频
		for i < n && clips[i][0] <= curEnd {
			nextEnd = algorithm.Max(nextEnd, clips[i][1])
			i++
		}
		// 找到下一个视频，更新 curEnd
		res++
		curEnd = nextEnd
		if curEnd >= T {
			// 已经可以拼出区间 [0, T]
			return res
		}
	}
	// 无法连续拼出区间 [0, T]
	return -1
}

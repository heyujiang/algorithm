package array

import "sort"

//2418.按身高排序  （https://leetcode.cn/problems/sort-the-people/）
func sortPeople(names []string, heights []int) []string {
	p := make([][]int, 0, len(heights))
	for i, v := range heights {
		p = append(p, []int{i, v})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i][1] > p[j][1]
	})

	res := make([]string, 0, len(heights))
	for _, v := range p {
		res = append(res, names[v[0]])
	}
	return res
}

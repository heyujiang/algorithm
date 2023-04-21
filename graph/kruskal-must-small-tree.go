package graph

import (
	"math"
	"sort"
)

//kruskal 最小生成树算法
//1135.最低成本联通所有城市
func minimumCost(n int, connections [][]int) int {
	uf := NewUF(n + 1)
	mst := 0

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	for _, v := range connections {
		if !uf.Connected(v[0], v[1]) {
			uf.Union(v[0], v[1])
			mst += v[2]
		}
	}

	if uf.Count() != 2 {
		return -1
	}
	return mst
}

//1584. 连接所有点的最小费用
func minCostConnectPoints(points [][]int) int {
	uf := NewUF(len(points))

	mst := 0
	pointCosts := make([][]int, 0)
	for i := 0; i < len(points); i++ {
		for j := 1; j < len(points); j++ {
			pointCosts = append(pointCosts, []int{i, j, int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))})
		}
	}

	sort.Slice(pointCosts, func(i, j int) bool {
		return pointCosts[i][2] < pointCosts[j][2]
	})

	for _, v := range pointCosts {
		if !uf.Connected(v[0], v[1]) {
			uf.Union(v[0], v[1])
			mst += v[2]
		}
	}
	return mst
}

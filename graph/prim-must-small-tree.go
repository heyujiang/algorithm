package graph

import (
	"container/heap"
	"math"
)

//Prim最小生成树算法
type Prim struct {
	gh        *graphHeap // 核心数据结构，存储横切边的优先级队列
	inMST     []bool     //累死visited数组的作用，记录哪些节点已经成为最小生成树的一部分
	weightSum int        //最小生成树的权重和
	graph     [][][]int  //用邻接表表示的一幅图 , graph[x] 记录节点x所有的相邻的边 ； 三元组[]int{from,to,weight} 表示一条边
}

func NewPrim(graph [][][]int) *Prim {
	gh := make(graphHeap, 0)
	heap.Init(&gh)

	p := &Prim{
		gh:        &gh,
		inMST:     make([]bool, len(graph)),
		weightSum: 0,
		graph:     graph,
	}

	p.inMST[0] = true
	p.cut(0)
	for p.gh.Len() > 0 {
		edge := heap.Pop(p.gh).([]int)

		if p.inMST[edge[1]] {
			continue
		}
		p.weightSum += edge[2]

		p.inMST[edge[1]] = true
		p.cut(edge[1])
	}

	return p
}

func (p *Prim) cut(s int) {
	for _, edge := range p.graph[s] {
		if p.inMST[edge[1]] {
			continue
		}
		heap.Push(p.gh, edge)
	}
}

//返回最小生成树的权重和
func (p *Prim) MinWeightSum() int {
	return p.weightSum
}

//判断最小生成树是否包含图中的所有节点
func (p *Prim) AllConnected() bool {
	for i := 0; i < len(p.graph); i++ {
		if !p.inMST[i] {
			return false
		}
	}
	return true
}

type graphHeap [][]int

func (g graphHeap) Len() int           { return len(g) }
func (g graphHeap) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g graphHeap) Less(i, j int) bool { return g[i][2] < g[j][2] }
func (g *graphHeap) Pop() interface{} {
	old := *g
	n := len(*g)
	x := old[n-1]
	*g = old[:n-1]
	return x
}
func (g *graphHeap) Push(x interface{}) {
	*g = append(*g, x.([]int))
}

//使用Prim最小生成树算法解决几道leetcode算法题
//1584. 连接所有点的最小费用
func minCostConnectPointsPrim(points [][]int) int {
	graph := make([][][]int, len(points))
	for i := 0; i < len(points); i++ {
		for j := 1; j < len(points); j++ {
			graph[i] = append(graph[i], []int{i, j, int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))})
		}
	}

	prim := NewPrim(graph)

	return prim.MinWeightSum()
}

//1135.最低成本联通所有城市
func minimumCostPrim(n int, connections [][]int) int {
	graph := make([][][]int, n)

	for _, connection := range connections {
		u := connection[0] - 1
		v := connection[1] - 1
		weight := connection[2]
		graph[u] = append(graph[u], []int{u, v, weight})
		graph[v] = append(graph[v], []int{v, u, weight})
	}

	prim := NewPrim(graph)
	if prim.AllConnected() {
		return prim.MinWeightSum()
	}
	return -1
}

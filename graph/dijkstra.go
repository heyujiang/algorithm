package graph

import (
	"container/heap"
	"math"
)

//获得节点到节点的权重
func weight(from, to int) int {
	return 0
}

//获得节点的邻节点
func adj(x int) []int {
	return []int{}
}

type DijkstraState struct {
	Id            int
	DistFromStart int
}

func dijkstra(start int, graph [][]int) []int {
	distTo := make([]int, len(graph))
	for i := range distTo {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0

	dq := make(dijkstraQueue, 0)
	heap.Init(&dq)
	heap.Push(&dq, DijkstraState{start, 0})

	for dq.Len() > 0 {
		curState := heap.Pop(&dq).(DijkstraState)
		curId := curState.Id
		curDistFromStart := curState.DistFromStart

		if curDistFromStart > distTo[curId] { //已经有一条更短的路径到达该节点
			continue
		}

		for _, nextId := range adj(curId) {
			distToNextId := distTo[curId] + weight(curId, nextId) //start节点到下一个节点的权重
			if distTo[nextId] > distToNextId {
				distTo[nextId] = distToNextId
				heap.Push(&dq, DijkstraState{Id: nextId, DistFromStart: distToNextId})
			}
		}
	}

	return distTo
}

/*
指定end的算法
func dijkstra(start int, end int , graph [][]int) int {
	distTo := make([]int,len(graph))
	for i := range distTo {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0

	dq := make(dijkstraQueue,0)
	heap.Init(&dq)
	heap.Push(&dq,DijkstraState{start,0})

	for dq.Len() > 0 {
		curState := heap.Pop(&dq).(DijkstraState)
		curId := curState.Id
		curDistFromStart := curState.DistFromStart
		if curId == end {   //直接返回
			return curDistFromStart
		}
		if curDistFromStart > distTo[curId] { //已经有一条更短的路径到达该节点
			continue
		}

		for _,nextId := range adj(curId) {
			distToNextId := distTo[curId] + weight(curId,nextId)  //start节点到下一个节点的权重
			if distTo[nextId] > distToNextId {
				distTo[nextId] = distToNextId
				heap.Push(&dq,DijkstraState{Id:nextId,DistFromStart:distToNextId})
			}
		}
	}

	return 0 //start不可到达end
}
*/

type dijkstraQueue []DijkstraState

func (d dijkstraQueue) Len() int {
	return len(d)
}
func (d dijkstraQueue) Less(i, j int) bool {
	return d[i].DistFromStart < d[j].DistFromStart
}
func (d dijkstraQueue) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d *dijkstraQueue) Pop() interface{} {
	old := *d
	n := len(old)
	v := old[n-1]
	*d = old[:n-1]
	return v
}
func (d *dijkstraQueue) Push(x interface{}) {
	*d = append(*d, x.(DijkstraState))
}

//使用dijkstra算法解决几道leetcode算法题
//743。网络延迟时间
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n)

	for _, time := range times {
		from := time[0] - 1
		to := time[1] - 1
		weight := time[2]

		graph[from] = append(graph[from], []int{from, to, weight})
	}

	start := k - 1
	distTo := make([]int, n)
	for i := range distTo {
		distTo[i] = math.MaxInt32
	}
	distTo[start] = 0

	dq := make(dijkstraQueue, 0)
	heap.Init(&dq)
	heap.Push(&dq, DijkstraState{start, 0})

	for dq.Len() > 0 {
		curNode := heap.Pop(&dq).(DijkstraState)
		curId := curNode.Id
		curDistFromStart := curNode.DistFromStart

		if distTo[curId] < curDistFromStart {
			continue
		}

		for _, v := range graph[curId] {
			nextNodeId := v[1]
			nextDistFromCur := v[2]

			nextDistFromStart := nextDistFromCur + curDistFromStart

			if distTo[nextNodeId] > nextDistFromStart {
				distTo[nextNodeId] = nextDistFromStart
				heap.Push(&dq, DijkstraState{Id: nextNodeId, DistFromStart: nextDistFromStart})
			}
		}
	}

	res := 0
	for _, dist := range distTo {
		if res < dist {
			res = dist
		}
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

//1631. 最小体力消耗路径
func minimumEffortPath(heights [][]int) int {
	return 0
}

//1514. 概率最大的路径
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := maxProbabilityBuildGraph(n, edges, succProb)

	toProbs := make([]float64, n)
	toProbs[start] = 1

	ph := make(ProbStateHeap, 0)
	heap.Init(&ph)
	heap.Push(&ph, ProbState{id: start, probFromStart: 1})

	for ph.Len() > 0 {
		curNode := heap.Pop(&ph).(ProbState)
		if curNode.id == end {
			return curNode.probFromStart
		}
		if toProbs[curNode.id] > curNode.probFromStart {
			continue
		}

		for _, nextNodeTuple := range graph[curNode.id] {
			curNodeId := nextNodeTuple.to
			nextProbFromCur := nextNodeTuple.prob
			nextProbFromStart := toProbs[curNode.id] * nextProbFromCur

			if nextProbFromStart > toProbs[curNodeId] {
				toProbs[curNodeId] = nextProbFromStart
				heap.Push(&ph, ProbState{curNodeId, nextProbFromStart})
			}
		}
	}

	return 0
}

type tuple struct {
	to   int
	prob float64
}
type ProbState struct {
	id            int
	probFromStart float64
}

type ProbStateHeap []ProbState

func (p ProbStateHeap) Len() int           { return len(p) }
func (p ProbStateHeap) Less(i, j int) bool { return p[i].probFromStart > p[j].probFromStart }
func (p ProbStateHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *ProbStateHeap) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}
func (p *ProbStateHeap) Push(x interface{}) {
	*p = append(*p, x.(ProbState))
}

func maxProbabilityBuildGraph(n int, edges [][]int, succProb []float64) [][]tuple {
	graph := make([][]tuple, n)
	for i, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], tuple{to: edge[1], prob: succProb[i]})
		graph[edge[1]] = append(graph[edge[1]], tuple{to: edge[0], prob: succProb[i]})
	}

	return graph
}

package graph

//797 所有可能路径  && 剑指 Offer II 110. 所有路径

var res [][]int
var onPath []int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0)
	onPath = make([]int, 0)
	traverse(graph, 0)
	return res
}

func traverse(graph [][]int, s int) {
	onPath = append(onPath, s)

	if s == len(graph)-1 {
		res = append(res, append([]int{}, onPath...))
	}

	for _, v := range graph[s] {
		traverse(graph, v)
	}

	onPath = onPath[:len(onPath)-1]
}

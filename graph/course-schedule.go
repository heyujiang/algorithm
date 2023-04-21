package graph

//207.课程表
//核心原理是判断图中是否有环
//DFS 解法
func canFinish(numCourses int, prerequisites [][]int) bool {
	onPath := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	hasCycle := false
	graph := buildGraph(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		if hasCycle { //已经找到环了 直接返回
			break
		}
		canFinishTraverse(graph, i, &onPath, &visited, &hasCycle)
	}
	return !hasCycle
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	return graph
}

func canFinishTraverse(graph [][]int, s int, onPath *[]bool, visited *[]bool, hasCycle *bool) {
	if (*onPath)[s] {
		*hasCycle = true
		return
	}

	if (*visited)[s] { //已经遍历过节点
		return
	}

	(*visited)[s] = true
	(*onPath)[s] = true
	for _, v := range graph[s] {
		canFinishTraverse(graph, v, onPath, visited, hasCycle)
	}

	(*onPath)[s] = false
}

//环检测算法 BFS解法，先构建节点入度数组
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		indegree[v[0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	graph := buildGraph(numCourses, prerequisites)

	count := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, v := range graph[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return count == numCourses
}

//210 . 课程表II
// 拓扑排序 DFS解法
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(numCourses, prerequisites)

	res := make([]int, 0, numCourses)
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false

	for i := 0; i < numCourses; i++ {
		if hasCycle {
			return []int{}
		}
		findOrderTraverse(graph, i, &res, &visited, &onPath, &hasCycle)
	}

	reverse(res)
	return res
}

func findOrderTraverse(graph [][]int, x int, res *[]int, visited *[]bool, onPath *[]bool, hasCycle *bool) {
	if (*onPath)[x] {
		*hasCycle = true
		return
	}

	if (*visited)[x] {
		return
	}

	(*visited)[x] = true
	(*onPath)[x] = true
	for _, v := range graph[x] {
		findOrderTraverse(graph, v, res, visited, onPath, hasCycle)
	}
	(*onPath)[x] = false

	*res = append(*res, x)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//拓扑排序 BFS解法
func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		indegree[v[0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	graph := buildGraph(numCourses, prerequisites)

	count := 0
	res := make([]int, 0, numCourses)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, v := range graph[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
		res = append(res, cur)
	}

	if count == numCourses {
		return res
	}
	return []int{}
}

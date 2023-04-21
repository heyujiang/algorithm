package graph

//785 判断二分图
//核心原理是给图中的节点染色，相邻的节点颜色不同即为二分表，否则不是二分表
//DFS解法
func isBipartite(graph [][]int) bool {
	color := make([]bool, len(graph)) //布尔值来表示两种不同的颜色
	visit := make([]bool, len(graph))
	isBipartite := true
	for i := 0; i < len(graph); i++ {
		if !isBipartite {
			break
		}
		isBipartiteTraverse(graph, i, &visit, &color, &isBipartite)
	}
	return isBipartite
}

func isBipartiteTraverse(graph [][]int, x int, visited *[]bool, color *[]bool, isBipartite *bool) {
	if !*isBipartite {
		return
	}

	(*visited)[x] = true
	for _, v := range graph[x] {
		if !(*visited)[v] {
			(*color)[v] = !(*color)[x]
			isBipartiteTraverse(graph, v, visited, color, isBipartite)
		} else {
			if (*color)[v] == (*color)[x] {
				*isBipartite = false
				return
			}
		}
	}
}

//BFS解法
func isBipartiteBFS(graph [][]int) bool {
	color := make([]bool, len(graph)) //布尔值来表示两种不同的颜色
	visit := make([]bool, len(graph))
	isBipartite := true
	for i := 0; i < len(graph); i++ {
		if !isBipartite {
			break
		}
		bfsTraverse(graph, i, &visit, &color, &isBipartite)
	}
	return isBipartite
}

func bfsTraverse(graph [][]int, x int, visited *[]bool, color *[]bool, isBipartite *bool) {
	queue := make([]int, 0)
	queue = append(queue, x)
	(*visited)[x] = true
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			cur := queue[0]
			queue = queue[1:]

			for _, v := range graph[cur] {
				if !(*visited)[v] {
					(*color)[v] = !(*color)[cur]
					queue = append(queue, v)
					(*visited)[v] = true
				} else {
					if (*color)[v] == (*color)[cur] {
						*isBipartite = false
						return
					}
				}
			}
		}
	}
}

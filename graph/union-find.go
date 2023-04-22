package graph

//并查集（union-find）算法
type UF struct {
	num    int
	parent []int
}

func NewUF(n int) *UF {
	uf := &UF{
		num:    n,
		parent: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (u *UF) Count() int {
	return u.num
}

func (u *UF) Union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP != rootQ {
		u.parent[rootP] = rootQ
		u.num--
	}
}

func (u *UF) Connected(p, q int) bool {
	rootP := u.find(p)
	rootQ := u.find(q)
	return rootP == rootQ
}

func (u *UF) find(x int) int {
	//for u.parent[x] != x {
	//	x = u.parent[x]
	//}
	//return x

	if u.parent[x] != x { //优化后 将一颗树压平只有两层
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

// 使用并查集算法解决leetcode算法题
//323 。 无向图中联通分量的数目
func countComponents(n int, edges [][]int) int {
	uf := NewUF(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	return uf.Count()
}

//990. 等式方程的可满足性
func equationsPossible(equations []string) bool {
	uf := NewUF(26)

	for _, v := range equations {
		if v[1] == '=' {
			uf.Union(int(v[0]-'a'), int(v[3]-'a'))
		}
	}

	for _, v := range equations {
		if v[1] == '!' {
			if uf.Connected(int(v[0]-'a'), int(v[3]-'a')) {
				return false
			}
		}
	}
	return true
}

//261 。 以图判断树
func validTree(n int, edges [][]int) bool {
	uf := NewUF(n)

	for _, v := range edges {
		if uf.Connected(v[0], v[1]) {
			return false
		}
		uf.Union(v[0], v[1])
	}

	return uf.Count() == 1
}

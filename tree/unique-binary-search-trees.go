package tree

//96. 不同的二叉搜索树
var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return numTreesTraverse(1, n)
}

func numTreesTraverse(lo, hi int) int {
	if lo > hi {
		return 1
	}

	if memo[lo][hi] >= 0 {
		return memo[lo][hi]
	}

	res := 0
	for i := lo; i <= hi; i++ {
		res += numTreesTraverse(lo, i-1) * numTreesTraverse(i+1, hi)
	}
	memo[lo][hi] = res
	return res
}

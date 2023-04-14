package array

import "fmt"

//54  旋转矩阵

func spiralOrder(matrix [][]int) []int {
	n ,m := len(matrix),len(matrix[0])
	res := make([]int,0,n*m)
	top,lower,left,right := 0,n-1,0,m-1
	for len(res) < n*m {
		if top <= lower {
			for i:=left;i<=right;i++ {
				res = append(res,matrix[top][i])
			}
			top++
		}

		if left <= right {
			for i:=top;i<=lower ;i++  {
				res = append(res,matrix[i][right])
			}
			right--
		}

		if top <= lower {
			for i := right; i >= left; i-- {
				res = append(res,matrix[lower][i])
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= top; i-- {
				res = append(res,matrix[i][left])
			}
			left++
		}
	}

	return res
}

// 59 螺旋矩阵II
func generateMatrix(n int) [][]int {
	res := make([][]int,n)
	for i:=0;i<n;i++{
		res[i] = make([]int,n)
	}
	top,lower,left,right := 0,n-1,0,n-1
	index := 0
	for index <= n*n {
		fmt.Println(index)
		if top <= lower {
			for i:=left;i<=right;i++ {
				res[top][i] = index
				index++
			}
			top++
		}

		if left <= right {
			for i:=top;i<=lower ;i++  {
				res[i][right] = index
				index++
			}
			right--
		}

		if top <= lower {
			for i := right; i >= left; i-- {
				res[lower][i] = index
				index++
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= top; i-- {
				res[i][left] = index
				index++
			}
			left++
		}
	}

	return res
}
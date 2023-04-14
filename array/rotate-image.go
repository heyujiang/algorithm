package array

func rotate(matrix[][]int)  {
	n := len(matrix)-1
	for i:=0;i<n;i++ {
		for j:=0;j<n-i;j++{
			matrix[i][j],matrix[n-j][n-i] = matrix[n-j][n-i],matrix[i][j]
		}
	}

	for i := 0; i <= n/2; i++ {
		for j:=0;j<= n ; j++ {
			matrix[i][j],matrix[n-i][j]  = matrix[n-i][j] ,matrix[i][j]
		}
	}
}

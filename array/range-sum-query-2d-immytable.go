package array

//304. 二维区域和检索 - 矩阵不可变   && 剑指 Offer II 013. 二维子矩阵的和
type NumMatrix struct {
	preMatrix [][]int
}


func ConstructorMatrix(matrix [][]int) NumMatrix {
	preMatrix := make([][]int,len(matrix)+1)
	for i:=0;i<len(preMatrix) ;i++  {
		preMatrix[i] = make([]int,len(matrix[0])+1)
	}

	for i := 0 ;i < len(matrix) ; i++  {
		for j := 0 ;j < len(matrix[0]) ; j++  {
			preMatrix[i+1][j+1] = matrix[i][j] + preMatrix[i+1][j] + preMatrix[i][j+1] - preMatrix[i][j]
		}
	}
	return NumMatrix{preMatrix:preMatrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preMatrix[row2+1][col2+1] + this.preMatrix[row1][col1] - this.preMatrix[row2+1][col1] - this.preMatrix[row1][col2+1]
}


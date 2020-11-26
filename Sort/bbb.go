package Sort

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxValue(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < row; i-- {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = Max(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[row-1][col-1]
}

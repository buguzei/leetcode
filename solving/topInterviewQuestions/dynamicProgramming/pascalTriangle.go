package dynamicProgramming

func PascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = make([]int, 1)
	triangle[0][0] = 1

	for i := 1; i < numRows; i++ {
		triangle[i] = make([]int, i+1)

		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
				continue
			}

			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}

	return triangle
}

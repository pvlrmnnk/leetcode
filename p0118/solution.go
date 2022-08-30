package p0118

type PascalsTriangleSolution func(numRows int) [][]int

func Solution(level int) [][]int {
	triangle := make([][]int, level)

	for i := 0; i < level; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

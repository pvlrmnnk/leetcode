package p0119

type PascalsTriangleSolution func(rowIndex int) []int

func Solution(level int) []int {
	row := make([]int, level+1)

	row[0] = 1
	for i := 0; i <= level; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

package p1996

import "sort"

type TheNumberofWeakCharactersSolution func(properties [][]int) int

// Time complexity: O(nlogn)
// Space complexity: O(1)
func Solution(properties [][]int) int {
	const (
		attack  = 0
		defence = 1
	)

	// sort by attack DESC, defence ASC
	sort.Slice(properties, func(i, j int) bool {
		switch {
		case properties[i][attack] > properties[j][attack]:
			return true
		case properties[i][attack] == properties[j][attack] && properties[i][defence] < properties[j][defence]:
			return true
		default:
			return false
		}
	})

	var weakCount, maxDefence int
	for _, characterProperties := range properties {
		if characterProperties[defence] < maxDefence {
			weakCount++
		} else {
			maxDefence = characterProperties[defence]
		}
	}

	return weakCount
}

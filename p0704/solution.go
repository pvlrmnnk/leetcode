package p0704

import "golang.org/x/exp/constraints"

type BinarySearchSolution func(nums []int, target int) int

func Solution(nums []int, target int) int {
	// also commom names are low, high, mid
	left, right, pivot := 0, len(nums)-1, 0

	for left <= right {
		// when a < b then a + (b - a) / 2 equals (a + b) / 2
		// this solution is overflow-proof
		pivot = left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return -1
}

func GenericsSolution[T constraints.Ordered](nums []T, target T) int {
	low, high, mid := 0, len(nums)-1, 0

	for low <= high {
		mid = low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

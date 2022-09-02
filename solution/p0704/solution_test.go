package p0704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   BinarySearchSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				nums   []int
				target int
			}
			tcs := []struct {
				name string
				args args
				want int
			}{
				{
					"target is in nums",
					args{[]int{-1, 0, 3, 5, 9, 12}, 9},
					4,
				},
				{
					"target is not in nums",
					args{[]int{-1, 0, 3, 5, 9, 12}, 2},
					-1,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.nums, tc.args.target)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

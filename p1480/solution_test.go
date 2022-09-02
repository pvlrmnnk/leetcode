package p1480

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSumOf1dArraySolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   RunningSumOf1dArraySolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				nums []int
			}
			tcs := []struct {
				name string
				args args
				want []int
			}{
				{
					"",
					args{[]int{1, 2, 3, 4}},
					[]int{1, 3, 6, 10},
				},
				{
					"",
					args{[]int{1, 1, 1, 1, 1}},
					[]int{1, 2, 3, 4, 5},
				},
				{
					"",
					args{[]int{3, 1, 2, 10, 1}},
					[]int{3, 4, 6, 16, 17},
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.nums)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

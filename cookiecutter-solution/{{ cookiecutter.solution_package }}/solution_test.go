package {{ cookiecutter.solution_package }}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test{{ cookiecutter.solution_fn_type }}(t *testing.T) {
	solutions := []struct {
		name string
		fn   {{ cookiecutter.solution_fn_type }}
	}{
		{"solution", Solution},
	}

	for _, solution := range solutions {
		t.Run(solution.name, func(t *testing.T) {
			testCases := []struct {
				name string
			}{
				{
					"",
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					//_ := solution.fn()
					assert.True(t, true)
				})
			}
		})
	}
}

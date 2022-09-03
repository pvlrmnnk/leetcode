package {{ cookiecutter.solution_package }}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test{{ cookiecutter.solution_fn_type }}(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   {{ cookiecutter.solution_fn_type }}
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				//
			}
			tcs := []struct {
				name string
				args args
				want any
			}{
				{
					"",
					args{},
					nil,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

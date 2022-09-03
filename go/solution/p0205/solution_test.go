package p0205

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsomorphicStringsSolution(t *testing.T) {
	t.Parallel()

	ss := []struct {
		name string
		fn   IsomorphicStringsSolution
	}{
		{"solution", Solution},
	}

	for _, s := range ss {
		s := s
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()

			type args struct {
				string1, string2 string
			}
			tcs := []struct {
				name string
				args args
				want bool
			}{
				{
					"",
					args{"egg", "add"},
					true,
				},
				{
					"",
					args{"foo", "bar"},
					false,
				},
				{
					"",
					args{"paper", "title"},
					true,
				},
			}

			for _, tc := range tcs {
				tc := tc
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					got := s.fn(tc.args.string1, tc.args.string2)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}

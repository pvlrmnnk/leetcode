package p0205

import "strings"

type IsomorphicStringsSolution func(s string, t string) bool

func Solution() IsomorphicStringsSolution {
	return func(s, t string) bool {
		transform := func(s string) string {
			var ms strings.Builder
			m := make(map[rune]rune)
			i := 0

			for _, r := range s {
				var nr rune
				if _, ok := m[r]; !ok {
					nr = rune(i)
					i++
					m[r] = nr
				} else {
					nr = m[r]
				}
				ms.WriteRune(nr)
			}

			return ms.String()
		}

		return transform(s) == transform(t)
	}
}

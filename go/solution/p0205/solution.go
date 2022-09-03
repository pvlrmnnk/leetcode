package p0205

import "strings"

type IsomorphicStringsSolution func(string1, string2 string) bool

func Solution(string1, string2 string) bool {
	//nolint:varnamelen
	transform := func(s string) string {
		var b strings.Builder
		dic := make(map[rune]rune)

		i := 0
		for _, r := range s {
			var nr rune
			if _, ok := dic[r]; !ok {
				nr = rune(i)
				i++
				dic[r] = nr
			} else {
				nr = dic[r]
			}
			b.WriteRune(nr)
		}

		return b.String()
	}

	return transform(string1) == transform(string2)
}

package p0392

type IsSubsequenceSolution func(s string, t string) bool

//nolint:varnamelen
func Solution(s, t string) bool {
	var i, j int

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return len(s) == i
}

package test

import "fmt"

func N(s string, n int) string {
	return fmt.Sprintf("%s #%d", s, n+1)
}

func Solution(n int) string {
	return N("Solution", n)
}

func TestCase(n int) string {
	return N("TestCase", n)
}

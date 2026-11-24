package main

import "fmt"

// ZArray returns z where z[i] is the length of the longest substring
// starting at i that matches a prefix of s (z[0] = len(s) by convention).
// The [l, r) window of the rightmost known prefix-match lets previously
// computed values seed new ones, keeping the whole pass linear.
func ZArray(s string) []int {
	n := len(s)
	z := make([]int, n)
	if n == 0 {
		return z
	}
	z[0] = n
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i] > r {
			l, r = i, i+z[i]
		}
	}
	return z
}

// Search returns all indexes where pattern occurs in text, by computing the
// Z-array of pattern + "\x00" + text: any z value equal to len(pattern)
// marks a match.
func Search(text, pattern string) []int {
	out := []int{}
	m := len(pattern)
	if m == 0 || m > len(text) {
		return out
	}
	z := ZArray(pattern + "\x00" + text)
	for i := m + 1; i < len(z); i++ {
		if z[i] == m {
			out = append(out, i-m-1)
		}
	}
	return out
}

func main() {
	fmt.Println(ZArray("aabxaab"))
	fmt.Println(Search("aabxaabxcaabxaabxay", "aabx"))
}

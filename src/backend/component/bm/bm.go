package bm

import (
	"strings"
)

const limit = 256

func badCharHeuristic(str string, size int, badchar [limit]int) {
	for i := 0; i < limit; i++ {
		badchar[i] = -1
	}
	for i := 0; i < size; i++ {
		badchar[int(str[i])] = i
	}
}

func BoyerMooreSearch(txt string, pat string) int64 {
	txt = strings.ToLower(txt)
	pat = strings.ToLower(pat)

	m := len(pat)
	n := len(txt)
	var badchar [limit]int

	badCharHeuristic(pat, m, badchar)

	s := 0
	for s <= (n - m) {
		j := m - 1

		for j >= 0 && pat[j] == txt[s+j] {
			j--
		}

		if j < 0 {
			result := int64(s)
			return result
		} else {
			if (j - badchar[txt[s+j]]) > 1 {
				s += j - badchar[txt[s+j]]
			} else {
				s++
			}
		}
	}
	return -1
}

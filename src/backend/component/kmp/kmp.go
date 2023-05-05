package kmp

import (
	"strings"
)

func computeLPSArray(pat string, M int, lps []int) {
	len := 0
	lps[0] = 0
	i := 1
	for i < M {
		if pat[i] == pat[len] {
			len++
			lps[i] = len
			i++
		} else {
			if len != 0 {
				len = lps[len-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

func KMPSearch(pat string, txt string) int64 {
	pat = strings.ToLower(pat)
	txt = strings.ToLower(txt)

	M := len(pat)
	N := len(txt)

	lps := make([]int, M)

	computeLPSArray(pat, M, lps)

	i, j := 0, 0
	for i < N {
		if pat[j] == txt[i] {
			j++
			i++
		}

		if j == M {
			result := int64(i - j)
			return result
		} else if i < N && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}
	return -1
}

// func main() {
// 	txt := "ABABDABACDABABCABAB"
// 	pat := "ABABCABAB"
// 	KMPSearch(pat, txt)
// }

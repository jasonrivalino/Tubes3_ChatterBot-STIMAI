package lcs

// Fungi untuk menghitung panjang LCS dengan DP (Dynamic Programming) secara bottom-up
func lcs(a, b string) int {
	m := len(a)
	n := len(b)
	L := make([][]int, m+1)
	for i := 0; i < len(L); i++ {
		L[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0
			} else if a[i-1] == b[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = max(L[i-1][j], L[i][j-1])
			}
		}
	}

	return L[m][n]
}

// Fungsi untuk mengembalikan nilai maksimum dari dua bilangan
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fungsi untuk menentukan kemiripan antara dua string
func Similarity(a, b string) float64 {
	l := lcs(a, b)
	if len(a) > len(b) {
		return float64(l) / float64(len(a))
	}
	return float64(l) / float64(len(b))
}

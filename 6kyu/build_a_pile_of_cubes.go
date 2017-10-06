package codewars

func FindNb(m int) int {
	for i := 1; i < m; i++ {
		sub := (i * ( i + 1) / 2) * (i * ( i + 1) / 2)
		if sub == m {
			return i
		} else if sub > m {
			return -1
		}
	}
	return -1
}

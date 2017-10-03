package codewars

func HasUniqueChar (str string) bool {
	m := make(map[rune]int)
	for _, value := range str {
		if _, ok := m[value]; ok {
			return false
		} else {
			m[value] = 1
		}
	}
	return true
}

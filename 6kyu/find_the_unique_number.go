package codewars

func FindUniq(arr []float32) float32 {
	seenMap := make(map[float32]int)
	for _, val := range arr {
		_, inMap := seenMap[val]
		if inMap {
			seenMap[val] += 1
		} else {
			seenMap[val] = 1
		}
	}
	for key, val := range seenMap {
		if val == 1 {
			return key
		}
	}
	return arr[0]
}

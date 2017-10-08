package codewars

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	posPeaks := PosPeaks{[]int{}, []int{}}
	for i := 1; i < len(array)-1; i++ {
		if array[i] == array[i+1] && array[i] > array[i-1] {
			allSame, counter, endPosition := true, 0, i
			for j := i; j < len(array)-1; j++ {
				if array[j] != array[i] {
					if j != len(array) {
						endPosition = j
					} else {
						allSame = false
					}
					break
				}
				counter++
			}
			if allSame && counter > 1 && array[endPosition] < array[i]{
				posPeaks.Pos = append(posPeaks.Pos, i)
				posPeaks.Peaks = append(posPeaks.Peaks, array[i])
			}
		}
		if array[i-1] < array[i] && array[i] > array[i+1] {
			posPeaks.Pos = append(posPeaks.Pos, i)
			posPeaks.Peaks = append(posPeaks.Peaks, array[i])
		}
	}
	return posPeaks
}

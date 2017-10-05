package codewars

import "math"

//import "math"

//import "fmt"

//import "fmt"

//import "fmt"

// https://www.codewars.com/kata/is-my-friend-cheating


func RemovNb(m uint64) [][2]uint64 {
	sum := m * (m+1) / 2
	resArr := [][2]uint64{}
	upperBound := uint64(math.Sqrt(float64(sum + 1) - 1))
	lowerBound := ((m - 1) * m / 2) / (m + 1)
	for i := upperBound; i >= lowerBound; i-- {
		currNum := (sum - i) / (i+1)
		if i * currNum + i + currNum == sum {
			resArr = append(resArr, [2]uint64{i, currNum})
			resArr = append(resArr, [2]uint64{currNum, i})
		}
	}
	if len(resArr) > 0 {
		return resArr
	} else {
		return nil
	}
}
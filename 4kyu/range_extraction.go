package codewars

import (
	"strconv"
	"math"
)

func Solution(list []int) string {
	counter, retString, tempRetStr, startVal, setZero := 0, "", "", 0, false
	for i := 0; i < len(list)-1; i++ {
		prod := math.Abs(float64(list[i+1]) -float64(list[i]))
		if prod == 1 {
			if startVal == 0 && setZero == false {
				startVal = list[i]
				if list[i] == 0 {
					setZero = true
				}
			}
			tempRetStr += strconv.Itoa(list[i])
			counter++
		} else {
			if counter >= 2 || counter <= -2 {
				retString += strconv.Itoa(startVal) + "-" + strconv.Itoa(list[i]) + ","
			} else {
				if tempRetStr != "" {
					retString += tempRetStr + ","
				}
				retString += strconv.Itoa(list[i]) + ","
			}
			tempRetStr = ""
			counter = 0
			startVal = 0
			setZero = false
		}
	}
	if tempRetStr != "" {
		if list[len(list)-1] - list[len(list)-2] == 1 && counter >= 2 {
			retString += strconv.Itoa(startVal) + "-" + strconv.Itoa(list[len(list)-1])
		} else {
			retString += tempRetStr + "," + strconv.Itoa(list[len(list)-1])
		}
	} else {
		retString += strconv.Itoa(list[len(list)-1])
	}
	return retString
}

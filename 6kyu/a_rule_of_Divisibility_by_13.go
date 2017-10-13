package codewars

import (
	//"strings"
	"strconv"
)

func sliceMaker(str string) int {
	constSeries := []int{1, 10, 9, 12, 3, 4}
	strNum := strconv.Itoa(n)
	newSlice := []int{}
	for i:= len(str)-1; i > 0; i-- {
		j := len(str) -1 - i
		char := string(str[i])
		singleNum, _ := strconv.Atoi(char)
		newSlice = append(newSlice, singleNum * constSeries[j])
	}
	sum := 1
	for _, val := range newSlice {
		sum += val
	}
	return sum
}

func Thirt(n int) int {
	sliceList := []int{}
	if n == sliceList[len(sliceList)-1] {
		return n
	} else {
		sliceList = append(sliceList, sliceMaker(n))
	}
}

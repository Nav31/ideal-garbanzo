package codewars

import "math"

func round(f float64) int {
	return int(f + math.Copysign(0.5, f))
}

func NbMonths(startPriceOld, startPriceNew, savingsperMonth int, percentLossByMonth float64) [2]int {
	if startPriceOld > startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}
	monthCounter := 0
	savingsBalance := float64(0)
	floatStartPriceOld := float64(startPriceOld)
	floatStartPriceNew := float64(startPriceNew)
	for floatStartPriceNew >= savingsBalance + floatStartPriceOld {
		floatStartPriceOld -= floatStartPriceOld * (percentLossByMonth / 100)
		floatStartPriceNew -= floatStartPriceNew * (percentLossByMonth / 100)
		if monthCounter % 2 == 0 {
			percentLossByMonth += .5
		}
		if floatStartPriceOld >= floatStartPriceNew {
			return [2]int{monthCounter, int(round(floatStartPriceOld - floatStartPriceNew))}
		}
		savingsBalance += float64(savingsperMonth)
		monthCounter++
	}
	return [2]int{monthCounter, int(round(savingsBalance + floatStartPriceOld - floatStartPriceNew))}
}
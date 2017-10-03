package codewars

//import "math"

func NbMonths(startPriceOld, startPriceNew, savingsperMonth int, percentLossByMonth float64) [2]int {
	monthCounter := 0
	// see when savings balance is greater than startPriceNew
	savingsBalance := float64(savingsperMonth)
	floatStartPriceOld := float64(startPriceOld)
	floatStartPriceNew := float64(startPriceNew)

	for floatStartPriceNew > savingsBalance {
		monthlyDeprecation := float64(0)
		monthlyDeprecation = float64(floatStartPriceOld) * float64(.0025) // deprecation at fixed rate
		floatStartPriceOld -= monthlyDeprecation
		floatStartPriceNew -= floatStartPriceNew * (percentLossByMonth / 100)
		savingsBalance += float64(savingsperMonth) - monthlyDeprecation
		if savingsBalance > floatStartPriceNew {
			return [2]int{monthCounter, int(savingsBalance - floatStartPriceNew)}
		}
		monthCounter++
	}
	return [2]int{monthCounter, int(savingsBalance - floatStartPriceNew)}
}
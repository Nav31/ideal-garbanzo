package codewars

//import "math"

func NbMonths(startPriceOld, startPriceNew, savingsperMonth int, percentLossByMonth float64) [2]int {
	monthCounter := 0
	// see when savings balance is greater than startPriceNew
	savingsBalance := savingsperMonth
	for startPriceNew > savingsBalance {
		monthlyDeprecation := 0
		//if monthCounter % 2 == 0 { // bi-monthly
		//	monthlyDeprecation = int(float64(startPriceOld) * float64(.025)) // deprecation at fixed rate
		//	startPriceOld -= monthlyDeprecation
		//}
		// monthly
		monthlyDeprecation = int(float64(startPriceOld) * float64(.025)) // deprecation at fixed rate
		startPriceOld -= monthlyDeprecation
		startPriceNew -= int(float64(startPriceNew) * (percentLossByMonth / 100))
		savingsBalance += savingsperMonth - monthlyDeprecation
		monthCounter++
	}
	return [2]int{monthCounter, savingsBalance - startPriceNew}
}
package codewars


func ValidParentheses(parens string) bool {
	counter := 0
	for _, val := range parens {
		if string(val) == "(" {
			counter++
		} else if string(val) == ")" {
			counter--
		}
		if counter < 0 {
			return false
		}
	}
	if counter != 0 {
		return false
	} else {
		return true
	}
}

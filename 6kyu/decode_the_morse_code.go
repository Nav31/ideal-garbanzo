package codewars

import (
	"strings"
)

// https://www.codewars.com/kata/decode-the-morse-code/train/go
// Have to findout what the map is called on codewars

func DecodeMorse(morseCode string) string {
	wordSlice := strings.Split(morseCode, " ")
	retStr := ""
	isPresentPrev := 0
	for i := 0; i <  len(wordSlice); i++ {
		mapVal, isPresent := MORSE_CODE[wordSlice[i]]
		if isPresent {
			retStr += mapVal
		} else {
			isPresentPrev += 1
			if isPresentPrev >= 2 {
				retStr += " "
				isPresentPrev = 0
			}
		}
	}
	return strings.Trim(retStr, " ")
}
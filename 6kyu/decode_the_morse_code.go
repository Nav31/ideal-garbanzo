package codewars

import (
	"strings"
)

// https://www.codewars.com/kata/decode-the-morse-code/train/go
// Have to findout what the map is called on codewars

func DecodeMorse(morseCode string) string {
	morseAlphabet := map[string]string {
		"..-.": "F", "-..-": "X",
		".--.": "P", "-": "T", "..---": "2",
		"....-": "4", "-----": "0", "--...": "7",
		"...-": "V", "-.-.": "C", ".": "E", ".---": "J",
		"---": "O", "-.-": "K", "----.": "9", "..": "I",
		".-..": "L", ".....": "5", "...--": "3", "-.--": "Y",
		"-....": "6", ".--": "W", "....": "H", "-.": "N", ".-.": "R",
		"-...": "B", "---..": "8", "--..": "Z", "-..": "D", "--.-": "Q",
		"--.": "G", "--": "M", "..-": "U", ".-": "A", "...": "S", ".----": "1",
		}
	wordSlice := strings.Split(morseCode, " ")
	retStr := ""
	isPresentPrev := 0
	for i := 0; i <  len(wordSlice); i++ {
		mapVal, isPresent := morseAlphabet[wordSlice[i]]
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
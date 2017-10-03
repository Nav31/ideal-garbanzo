package codewars

import "strings"

func ToNato(words string) string {

	nato := map[string]string {
		"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta", "E": "Echo",
		"F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India", "J": "Juliett",
		"K": "Kilo","L": "Lima", "M": "Mike", "N": "November", "O": "Oscar",
		"P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango",
		"U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "X-ray",
		"Y": "Yankee", "Z": "Zulu",
		}

	retStr := ""
	for _, val := range words {
		natoCode, present := nato[strings.ToUpper(string(val))]
		if present {
			retStr += natoCode + " "
		} else {
			if string(val) != " " {
				retStr += string(val)
			}
		}
	}
	return strings.TrimSpace(retStr)
}

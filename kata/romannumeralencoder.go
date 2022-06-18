package kata

//nolint:gomnd
// RomanNumeralEncoder returns the Roman Numeral of integer N as a string
func RomanNumeralEncoder(i int) string {
	iRemaining := i
	romanNumeral := ""

	if iRemaining >= 100 {
		iRemaining %= 100
		romanNumeral += "C"
	}

	if iRemaining >= 90 {
		iRemaining %= 90
		romanNumeral += "XC"
	}

	if iRemaining >= 50 {
		iRemaining %= 50
		romanNumeral += "L"
	}

	if iRemaining >= 40 {
		iRemaining %= 40
		romanNumeral += "XL"
	}

	if iRemaining >= 10 {
		for ; iRemaining >= 10; iRemaining -= 10 {
			romanNumeral += "X"
		}

		iRemaining %= 10
	}

	if iRemaining >= 9 {
		iRemaining -= 9
		romanNumeral += "IX"
	}

	if iRemaining >= 5 {
		iRemaining %= 5
		romanNumeral += "V"
	}

	if iRemaining >= 4 {
		iRemaining -= 4
		romanNumeral += "IV"
	}

	for ; iRemaining >= 1; iRemaining-- {
		romanNumeral += "I"
	}

	return romanNumeral
}

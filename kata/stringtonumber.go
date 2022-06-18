package kata

import "strconv"

// StringToNumber converts a string representation of an integer to an integer
func StringToNumber(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

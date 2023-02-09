package kata

// EvenOrOdd that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
func EvenOrOdd(number int) string {
	if number%2 != 0 {
		return "Odd"
	}

	return "Even"
}

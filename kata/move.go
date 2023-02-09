package kata

// Move takes the current position of the hero and the roll (1-6) and return the new position.
//
// In this imaginary terminal game, the hero moves from left to right. The player rolls the dice and moves the number
// of spaces indicated by the dice two times.
func Move(position int, roll int) int {
	return position + roll*2
}

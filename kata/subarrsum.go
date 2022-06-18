package kata

// SubArrSum returns the maximum sum subarray of any sequence of consecutive numbers in an array or list of integers.
//
// When the list is made up of only positive numbers then the maximum sum is simply the sum of the whole array. If the
// list is made up of only negative numbers, return 0 instead. An empty list is considered to have zero greatest sum.
// Note that the empty list or array is also a valid sublist/subarray. If a list contains positive and negative numbers
// then you will need to work out which subsequence can be added together to give the highest total.
// (note: this may include only a small section of the numbers in the initial array)
func SubArrSum(numbers []int) int {
	maxSum := 0

	for i := 0; i < len(numbers); i++ {
		loopSum := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			loopSum += numbers[j]

			if loopSum > maxSum {
				maxSum = loopSum
			}
		}
	}

	if maxSum > 0 {
		return maxSum
	}

	return 0
}

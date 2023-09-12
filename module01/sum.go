package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// Iterative Solution
	res := 0 // tracking results here
	for _, x := range numbers {
		res += x
	}
	return res

	// Recursive solution
	// if len(numbers) == 0 {
	// 	return 0
	// }
	// return numbers[0] + Sum(numbers[1:])
}

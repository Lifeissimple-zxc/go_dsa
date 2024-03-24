package array

// findMax locates max value in an array
// Time Complexity is O(n)
func findMax(arr []int) int {
	res := arr[0]
	for _, val := range arr {
		if val <= res {
			continue
		}
		res = val
	}
	return res
}

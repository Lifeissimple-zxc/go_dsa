package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	// O(n), should be linear search
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

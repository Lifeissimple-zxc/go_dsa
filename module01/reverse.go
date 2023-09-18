package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// option1: linear solution
	// res := ""
	// for i := len(word)-1; i >= 0; i-- {
	// 	res += string(word[i])

	// }
	// return res

	// Option 2: O(n/2), but struggling to understand how to make it work for unicode
	// Get # of iterations
	// wLen := len(word) // saving to a variable bc of multiple uses
	// iters := wLen / 2
	// container := make([]byte, wLen)
	// for i := 0; i < iters; i++ {
	// 	right := wLen - i - 1
	// 	container[i], container[right] = word[right], word[i]
	// }
	// // Adding middle letter for words of odd len
	// if wLen%2 != 0 {
	// 	container[iters] = word[iters]
	// }
	// return string(container)

	// Option 3: Linear, but works with Unicode
	res := ""
	for _, ch := range word {
		res = string(ch) + res
	}
	return res

}

package main

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// option1: linear solution
	// res := ""
	// for i := len(word) - 1; i >= 0; i-- {
	// 	res += string(word[i])

	// }
	// return res

	// Get # of iterations
	var iters int
	wLen := len(word) // saving to a variable bc of multiple uses
	switch wLen % 2 {
	case 0:
		iters = wLen / 2
	default:
		iters = (wLen - 1) / 2
	}

	container := make([]byte, wLen)
	if iters%2 != 0 {
		container[iters+1] = word[iters+1]
	}
	for i := 0; i < iters; i++ {
		right := wLen - i - 1
		container[i] = word[right]
		container[right] = word[i]
	}

	return string(container)

}

func main() {
	fmt.Println(Reverse("hello"))
}

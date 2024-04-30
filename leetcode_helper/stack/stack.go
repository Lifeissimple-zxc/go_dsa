package stack

// TODO migrate this to generics

// Stack represents a stack for ints
type Stack struct {
	items []int
}

// Push adds an item on top of a stack
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop removes an item from the top of a stack and returns it.
// In case there are no elements in the stack, returns a zero.
func (s *Stack) Pop() (res int) {
	ln := len(s.items)
	if ln == 0 {
		return
	}
	res = s.items[ln-1]
	s.items = s.items[:ln-1]
	return res
}

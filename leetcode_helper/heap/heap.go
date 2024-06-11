package heap

import "errors"

// MaxHeap implements max heap data structure for ints using a slice.
type MaxHeapInt struct {
	s []int
}

// Insert adds an element to the heap
func (h *MaxHeapInt) Insert(val int) {
	// add to s as is
	h.s = append(h.s, val)
	// heapify up to maintain the heap property
	h.heapifyUp(len(h.s) - 1)

}

// Extract removes the larget element from the heap and returns it
func (h *MaxHeapInt) Extract() (int, error) {
	// handle over capacity extraction
	if len(h.s) == 0 {
		return 0, errors.New("heap does not have data, cannot extract")
	}
	root := h.s[0] // in max heap, root or zero index always contains the max

	// we need to rearrange the heap to maintain the heap property
	l := len(h.s) - 1
	h.s[0] = h.s[l]
	// shrink the slice
	h.s = h.s[:l]
	// maintaining heap property
	h.heapifyDown(0)

	return root, nil
}

func (h *MaxHeapInt) heapifyUp(i int) {
	// // swap if value under i is larger than parent

	// iterative implementation
	// for h.s[GetParent(i)] < h.s[i] {
	// 	h.swap(GetParent(i), i)
	// 	i = GetParent(i) // reassignment makes the loop work
	// }

	// recursive implementation
	// if parent is greater than element to swap, heap propery is ok so we can exit
	if h.s[GetParent(i)] >= h.s[i] {
		return
	}
	// no base case: perform swap and make a recursive call
	h.swap(GetParent(i), i)
	h.heapifyUp(GetParent(i))

}

func (h *MaxHeapInt) heapifyDown(i int) {
	lastIndex := len(h.s) - 1
	left, right := GetLeftChild(i), GetRightChild(i)
	childToCompare := 0

	for left <= lastIndex {
		if left == lastIndex {
			// left is the only child
			childToCompare = left
		} else if h.s[left] > h.s[right] {
			// left is larger
			childToCompare = left
		} else {
			// right is larger
			childToCompare = right
		}

		// swap if possible
		if h.s[i] < h.s[childToCompare] {
			h.swap(i, childToCompare)
			// update helper vars to continue traversal
			i = childToCompare
			left, right = GetLeftChild(i), GetRightChild(i)
		} else {
			return
		}

	}

}

func (h *MaxHeapInt) swap(x, y int) {
	h.s[x], h.s[y] = h.s[y], h.s[x]
}

func GetParent(i int) int {
	return (i - 1) / 2
}

func GetLeftChild(i int) int {
	return i*2 + 1
}

func GetRightChild(i int) int {
	return i*2 + 2
}

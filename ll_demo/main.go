package main

import (
	"fmt"
	"ll_demo/linkedlist"
)

func main() {
	// Reversal
	ll := linkedlist.LinkedList{}
	ll.Add(1, 2, 3, 4, 5)

	fmt.Println("List after creation", ll)

	ll.Reverse()
	fmt.Println("List after reversal", ll)

	ll.Reverse2()
	fmt.Println("List after reversal2", ll)

	ll.ReverseRecursive()
	fmt.Println("List after recursive reversal", ll)

	// Merging 2 sorted LLs
	llToMerge := &linkedlist.LinkedList{}
	fmt.Println("First list to merge", llToMerge)
	llToMerge2 := &linkedlist.LinkedList{}
	llToMerge2.Add(1, 3, 4)
	fmt.Println("Second list to merge", llToMerge2)

	mergedList := linkedlist.MergeSortedLists(llToMerge, llToMerge2)
	fmt.Println("Merged list:", mergedList)

}

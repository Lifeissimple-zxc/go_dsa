package main

import (
	"fmt"
	"ll_demo/linkedlist"
)

func main() {
	// Reversal
	fmt.Println("#### Reversing Linked List ####")
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
	fmt.Println("#### Merging Sorted Linked Lists ####")
	llToMerge := &linkedlist.LinkedList{}
	llToMerge.Add(2, 3)
	fmt.Println("First list to merge", llToMerge)
	llToMerge2 := &linkedlist.LinkedList{}
	llToMerge2.Add(1, 5, 7)
	fmt.Println("Second list to merge", llToMerge2)

	mergedList := linkedlist.MergeSortedLists(llToMerge, llToMerge2)
	fmt.Println("Merged list:", mergedList)

	// Merging using heads only
	mergedHead := linkedlist.MergeSortedListsUsingHeads(llToMerge.Head, llToMerge2.Head)
	mergedList2 := &linkedlist.LinkedList{Head: mergedHead}
	fmt.Println("LL Merged using Heads only:", mergedList2)

}

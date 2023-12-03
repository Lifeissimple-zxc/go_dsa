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
	fmt.Printf("Len of list: %d\n", ll.Len)
	fmt.Printf("List head: %v | List Tail %v\n", ll.Head, ll.Tail)

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
	llToMerge.Add(1, 2, 4)
	fmt.Println("First list to merge", llToMerge)
	llToMerge2 := &linkedlist.LinkedList{}
	llToMerge2.Add(1, 3, 4)
	fmt.Println("Second list to merge", llToMerge2)

	mergedList := linkedlist.MergeSortedLists(llToMerge, llToMerge2)
	fmt.Println("Merged list:", mergedList)

	// Merging using heads only
	fmt.Println("#### Merging Sorted Linked Lists w/o List Struct ####")
	mergedHead := linkedlist.MergeSortedListsUsingHeads(llToMerge.Head, llToMerge2.Head)
	mergedList2 := &linkedlist.LinkedList{Head: mergedHead}
	fmt.Println("LL Merged using Heads only:", mergedList2)

	// Finding intersection of two LLs
	fmt.Println("#### Finding intersection of two LLs ####")
	intersection := linkedlist.LinkedList{}
	intersection.Add(8, 4, 5)
	fmt.Println("Intersection: ", intersection)

	iLL1 := linkedlist.LinkedList{}
	iLL1.Add(4, 1)
	iLL1.Tail.Next = intersection.Head
	fmt.Println("LL1: ", iLL1)

	iLL2 := linkedlist.LinkedList{}
	iLL2.Add(5, 6, 1)
	iLL2.Tail.Next = intersection.Head
	fmt.Println("LL2: ", iLL2)

}

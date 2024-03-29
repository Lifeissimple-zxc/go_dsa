package main

import (
	"fmt"
	"leetcode_helper/array"
	"leetcode_helper/linkedlist"
)

func main() {
	// linked_list_exercises()
	array_exercises()
}

func linked_list_exercises() {
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

	intersectNode := linkedlist.FindIntersection(&iLL1, &iLL2)
	fmt.Println("Intersection search result (assuming list are unique):", intersectNode)

	intersectNodeDupl := linkedlist.FindIntersection(&iLL1, &iLL2)
	fmt.Println("Intersection search result ", intersectNodeDupl)

	intersectHeads := linkedlist.FindIntersectionHeads(iLL1.Head, iLL2.Head)
	fmt.Println("Intersection search result with heads only", intersectHeads)

	intersectHeadsDupl := linkedlist.FindIntersectionHeads(iLL1.Head, iLL2.Head)
	fmt.Println("Intersection search result with heads only", intersectHeadsDupl)

	fmt.Println("#### Checking if LL is a palindrome ####")
	palindrome := linkedlist.LinkedList{}
	palindrome.Add(1, 2, 2, 1)

	noPalindrome := linkedlist.LinkedList{}
	noPalindrome.Add(1, 2, 3)

	fmt.Printf(
		"Simple head-only palindrome check of a palindrome list  %s: %t\n",
		palindrome,
		linkedlist.SimplePalindromeCheckHead(palindrome.Head),
	)

	fmt.Printf(
		"Slow & Fast pointer palindrome check of a palindrome list  %s: %t\n",
		palindrome,
		linkedlist.IsPalindrome(&palindrome),
	)

	fmt.Printf(
		"Simple head-only Palindrome check of a non-palindrome list %s: %t\n",
		noPalindrome,
		linkedlist.SimplePalindromeCheckHead(noPalindrome.Head),
	)

	fmt.Printf(
		"Slow & Fast pointer palindrome check of a non-palindrome list %s: %t\n",
		noPalindrome,
		linkedlist.IsPalindrome(&noPalindrome),
	)

	fmt.Println("#### Checking if LL is has a cycle ####")
	nonCycledOneNodeLL := linkedlist.LinkedList{}
	nonCycledOneNodeLL.Add(1)
	fmt.Printf(
		"Does %s have a cycle? Result: %t\n",
		nonCycledOneNodeLL,
		linkedlist.HasCycleTwoPointers(&nonCycledOneNodeLL),
	)

	noCycleLL := linkedlist.LinkedList{}
	noCycleLL.Add(1, 2, 3)
	fmt.Printf(
		"Does %s have a cycle? Result: %t\n",
		noCycleLL,
		linkedlist.HasCycleTwoPointers(&noCycleLL),
	)

	LLWithCycle := linkedlist.LinkedList{}
	LLWithCycle.Add(3, 2)
	tl := LLWithCycle.Tail
	LLWithCycle.Add(0, 4)
	fmt.Printf("Pre cycling LL that is to be cycled now: %s\n", LLWithCycle)
	LLWithCycle.Tail.Next = tl

	fmt.Printf(
		"Does cycled list have a cycle? Result: %t\n",
		linkedlist.HasCycle(&LLWithCycle),
	)

	fmt.Printf(
		"Head only check for cycle for %s: %t\n",
		nonCycledOneNodeLL,
		linkedlist.HasCycleTwoPointersHead(nonCycledOneNodeLL.Head),
	)

	fmt.Printf(
		"Head only check for cycle for %s: %t\n",
		noCycleLL,
		linkedlist.HasCycleTwoPointersHead(noCycleLL.Head),
	)

	fmt.Printf(
		"Head only check for cycle for a list with cycle: %t\n",
		linkedlist.HasCycleTwoPointersHead(LLWithCycle.Head),
	)
}

func array_exercises() {
	fmt.Println("#### Finding Product within an array ####")
	nums := []int{3, 4, 5, 2}
	fmt.Printf("For array %+v max product is %d\n", nums, array.SimpleMaxProduct((nums)))
	nums2 := []int{1, 5, 4, 5}
	fmt.Printf("For array %+v max product is %d\n", nums2, array.SimpleMaxProduct((nums2)))
	nums3 := []int{3, 7}
	fmt.Printf("For array %+v max product is %d\n", nums3, array.SimpleMaxProduct((nums3)))
	nums4 := []int{10, 2, 5, 2}
	fmt.Printf("For array %+v max product is %d\n", nums4, array.SimpleMaxProduct((nums4)))

	fmt.Println("#### Building Pascal Triangle ####")
	fmt.Printf("Pascal triangle for %d is %+v\n", 5, array.PascalTriangle(5))

	fmt.Println("#### Locating single number ###")
	singleNumsArr := []int{2, 2, 1}
	singleNumsArr2 := []int{4, 1, 2, 1, 2}
	fmt.Printf(
		"%+v array's single number is %d\n",
		singleNumsArr,
		array.SingleNumber((singleNumsArr)),
	)
	fmt.Printf(
		"%+v array's single number is %d\n",
		singleNumsArr2,
		array.SingleNumber((singleNumsArr2)),
	)

	fmt.Printf(
		"%+v array's single number (bitwise XOR) is %d\n",
		singleNumsArr,
		array.SingleNumberBitWiseXOR((singleNumsArr)),
	)
	fmt.Printf(
		"%+v array's single number (bitwise XOR) is %d\n",
		singleNumsArr2,
		array.SingleNumberBitWiseXOR((singleNumsArr2)),
	)

	fmt.Println("#### Locating Major Number ###")
	majorNums1 := []int{3, 2, 3}
	majorNums2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf(
		"%+v array's major number is %d\n",
		majorNums1,
		array.MajEl((majorNums1)),
	)
	fmt.Printf(
		"%+v array's major number is %d\n",
		majorNums2,
		array.MajEl((majorNums2)),
	)
	fmt.Printf(
		"%+v array's major number (Boyer-Moore Majority Vote Algorithm) is %d\n",
		majorNums1,
		array.MajorElBoyerMoore((majorNums1)),
	)
	fmt.Printf(
		"%+v array's major number (Boyer-Moore Majority Vote Algorithm) is %d\n",
		majorNums2,
		array.MajorElBoyerMoore((majorNums2)),
	)

	fmt.Println("#### Moving Zeroes ###")
	moveZeroArr := []int{0, 1, 0, 3, 12}
	fmt.Printf("Array before MoveZeroes %+v\n", moveZeroArr)
	array.SimpleMoveZeroes(moveZeroArr)
	fmt.Printf("Array after MoveZeroes %+v\n", moveZeroArr)
	moveZeroArr2 := []int{1, 0, 0}
	fmt.Printf("Array2 before MoveZeroes %+v\n", moveZeroArr2)
	array.SimpleMoveZeroes(moveZeroArr2)
	fmt.Printf("Array2 after MoveZeroes %+v\n", moveZeroArr2)
	moveZeroArr3 := []int{2, 1}
	fmt.Printf("Array2 before MoveZeroes %+v\n", moveZeroArr3)
	array.SimpleMoveZeroes(moveZeroArr2)
	fmt.Printf("Array2 after MoveZeroes %+v\n", moveZeroArr3)

	fmt.Println("#### reverse vowels ###")
	res := array.ReverseVowels(" ipG0A4mAas::si0m4a0Gp0")
	fmt.Println(res)

	fmt.Println("#### Greatest Common Divisor of Strings ###")
	fmt.Println(array.GCDOfStrings("ABCABC", "ABC"))
}

package array

import (
	"fmt"
)

// TODO
func MaxProduct(nums []int) int {
	// max1 & max2 store max values
	// maxIx stores indeces to max values in the array
	var max1, max2, maxIx1, maxIx2 int
	for cur1, cur2 := 0, -1; cur2 < len(nums); {
		// first max
		if nums[cur1] > max1 {
			max1 = nums[cur1]
			maxIx1 = cur1
		}
		// move first cursor if possible
		if cur1+1 <= len(nums)-1 {
			cur1++
		}
		// [2,2,1,8,1,5,4,5,2,10,3,6,5,2,3] fails now
		if cur2 >= 0 && nums[cur2] > max2 && cur2 != maxIx1 {
			max2 = nums[cur2]
			maxIx2 = cur2
		}
		cur2++
		fmt.Println(nums[maxIx1], nums[maxIx2])
	}
	return (nums[maxIx1] - 1) * (nums[maxIx2] - 1)

}

// SimpleMaxProduct is a naive solution to maxProduct array problem
func SimpleMaxProduct(nums []int) int {
	var max, maxIx1, maxIx2 int
	// find max index 1
	for i, num := range nums {
		if num > max {
			maxIx1, max = i, num
		}
	}
	max = 0

	// find max index 2
	for i, num := range nums {
		// account for ix already occupied by first max
		if i == maxIx1 {
			continue
		}
		if num > max && num <= nums[maxIx1] {
			maxIx2, max = i, num
		}
	}
	// return product of the twop
	return (nums[maxIx1] - 1) * (nums[maxIx2] - 1)

}

func PascalTriangle(numRows int) [][]int {
	// Allocate array
	triangle := make([][]int, numRows)
	// Add triangle rows
	for i := range triangle {
		triangle[i] = generatePascalRow(i, triangle)
	}
	return triangle

}

func generatePascalRow(numRow int, triangle [][]int) []int {
	row := make([]int, numRow+1)
	for i := range row {
		if i == 0 || i == numRow {
			row[i] = 1
			continue
		}
		row[i] = triangle[numRow-1][i-1] + triangle[numRow-1][i]
	}
	return row
}

func SingleNumber(nums []int) int {
	// Keep track of what was seen how many times
	tracker := make(map[int]int)
	var res *int
	// Iterate over array & count occurrences
	for _, num := range nums {
		seen, ok := tracker[num]
		// Non-first time occurrence
		if ok {
			seen++
		} else {
			seen = 1
		}
		tracker[num] = seen
	}
	// Locate value with 1 occurrence
	for num, seen := range tracker {
		if seen == 1 {
			res = &num
			break
		}
	}
	return *res
}

func SingleNumberBitWiseXOR(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}

// MajEl locates element of the array that has >len(arr)
// occurences within it
func MajEl(nums []int) int {
	arrLen := len(nums)
	var res int
	// Edge case
	if arrLen == 1 {
		return nums[0]
	}
	// Count occurrences
	tracker := make(map[int]int)
	for _, num := range nums {
		seen, ok := tracker[num]
		if ok {
			seen++
		} else {
			seen = 1
		}
		tracker[num] = seen
	}
	// Return major element
	for num, seen := range tracker {
		if seen > (arrLen / 2) {
			res = num
			break
		}
	}
	return res
}

// MajorElBoyerMoore implements Boyer-Moore Majority Vote Algorithm
func MajorElBoyerMoore(nums []int) int {
	majority, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			majority, cnt = num, 1
			continue
		}
		if majority == num {
			cnt++
		} else {
			cnt--
		}
		if cnt > (len(nums) / 2) {
			break
		}
	}
	return majority
}

// SimpleMoveZeroes implements a naive solution to Move Zeroes problem
func SimpleMoveZeroes(nums []int) {
	// Traverse all elements of the array
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}
		// Swap with first non-zero
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}

	}
}

// BubbleSort implements a bubble sort algo: O(n^2)
func BubbleSort(nums []int) {
	// Iterate over array
	for i := 0; i < len(nums); i++ {
		// Inner loop over all elements but the last one
		// It decreases by i because outer loop shrinks # of elements to iterate over
		for j := 0; j < len(nums)-i-1; j++ {
			// If current is greater than next, we swap them
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// InsertionSort implements an insertion sort algo: O(n^2)
func InsertionSort(nums []int) {
	// Traverse the array starting at the 2nd element
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		lag_index := i - 1
		for lag_index > -1 && cur < nums[lag_index] {
			nums[lag_index+1] = nums[lag_index]
			lag_index--
		}
		nums[lag_index+1] = cur
	}
}

// SelectionSort implements an selection sort algo: O(n^2)
func SelectionSort(nums []int) {
	// find smallest element, put it to the start of the array
	for i := 0; i < len(nums); i++ {
		minIx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIx] {
				minIx = j
			}
		}
		if minIx != i {
			nums[i], nums[minIx] = nums[minIx], nums[i]
		}
	}
}

// MergeSort implements a merge sort algo does not work
func MergeSort(nums []int) []int {
	// Divide an conquer algo, recursive
	// base case
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left, right := nums[:mid], nums[mid:]
	return merge(MergeSort(left), MergeSort(right))

}

// merge is a building block for MergeSort
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			res = append(res, left[l])
		} else {
			res = append(res, right[r])
		}
	}
	// Account for difference in left and right lens
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res

}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	// base case: an already sorted array
	if start >= end {
		return
	}
	// choose pivot
	p := nums[medianThree(nums, start, end)]
	// partition such that all items to the left are <= pivot and all items to the right are >= pivot
	lp := partition(nums, start, end, p)
	// at this point we swapped all we can, we only need to swap left pointer with the pivot
	nums[lp], nums[end] = nums[end], nums[lp] // the last step of the partition!
	// sort pre and post partition parts of the inpuit
	quickSort(nums, start, lp-1)
	quickSort(nums, lp+1, end)
}

func partition(nums []int, start, end, pivot int) int {
	// allocate pointers for swaps that make the partitioning work
	lp, rp := start, end
	for lp < rp {
		// find first item from the left that is <= pivot
		for nums[lp] <= pivot && lp < rp {
			lp++
		}
		// find first item from the left that is >= pivot
		for nums[rp] >= pivot && lp < rp {
			rp-- // we are moving right to left so we decrement!s
		}
		// at this point we found items to swap
		nums[lp], nums[rp] = nums[rp], nums[lp]
	}
	return lp // we need lp hence returning it here
}

// medianThree chooses pivot index using median of three method.
func medianThree(nums []int, start, end int) int {
	mid := (end - start) / 2
	a, b, c := nums[start], nums[mid], nums[end]
	// is a our mid? != with two bool expressions is a XOR so one of the expressions needs to be false
	if (a > b) != (a > c) {
		// return start if it's the middle
		return start
	} else if (b > a) != (b > c) {
		return mid
	} else {
		return end
	}
}

package array

import "fmt"

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

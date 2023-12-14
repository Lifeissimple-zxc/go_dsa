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
	var max1, max2, maxIx1, maxIx2 int
	// find max index 1
	for i, num := range nums {
		if num > max1 {
			maxIx1, max1 = i, num
		}
	}

	// find max index 2
	for i, num := range nums {
		// account for ix already occupied by first max
		if i == maxIx1 {
			continue
		}
		if num > max2 && num <= max1 {
			maxIx2, max2 = i, num
		}
	}
	// return product of the twop
	return (nums[maxIx1] - 1) * (nums[maxIx2] - 1)

}
